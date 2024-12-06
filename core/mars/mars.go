package mars

import (
	"fmt"
	"math"
	"time"
	"math/rand"
)

func NewPlanet() *Planet {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	return &Planet{
		Temperature:    BaseTemp,
		WindSpeed:      0.0,
		Pressure:       BasePressure,
		Dust:           BaseDust,
		SunRadiation:   BaseSun,
		UVRadiation:    BaseUV,
		Seismicity:     NoSeismicity,
		IsDustStorm:    false,
		StormDuration:  0,
		StormIntensity: 1.0,
		Time: MartianTime{
			Sol:  0,
			Hour: 0.0,
			Min:  0.0,
		},
		rng: rng,
	}
}

func (p *Planet) Update() {
	p.advanceTime()
	p.updateConditions()
}

func (p *Planet) advanceTime() {
	p.Time.Min += 1.0
	if p.Time.Min >= MinsInHour {
		p.Time.Min -= MinsInHour
		p.Time.Hour += 1.0
	}
	if p.Time.Hour >= HoursInSol {
		p.Time.Hour -= HoursInSol
		p.Time.Sol += 1
	}
}

func (p *Planet) getLs() float64 {
	fraction := float64(p.Time.Sol % SolsInYear) / float64(SolsInYear)
	return fraction * 360.0
}

func (p *Planet) updateConditions() {
	ls := p.getLs()
	lsRad := ls * math.Pi / 180.0

	// Diurnal temperature variation
	dayFraction := (p.Time.Hour + p.Time.Min/MinsInHour) / HoursInSol
	dailyTemp := -70 + 80*math.Sin(2*math.Pi*dayFraction)

	// Seasonal temperature variation
	seasonalTemp := TempAmplitude * math.Sin(lsRad)

	p.Temperature = dailyTemp + seasonalTemp + p.randomVariation(-2, 2)

	// Pressure adjustment
	p.Pressure = BasePressure + PressureAmp*math.Cos(lsRad) + p.randomVariation(-0.05, 0.05)

	// Solar and UV radiation
	baseSun := BaseSun + SunAmp*math.Cos((ls-250.0)*math.Pi/180.0)
	dustFactor := math.Max(0.0, 1.0-p.Dust/MaxDust) // Ensure factor stays >= 0
	p.SunRadiation = math.Max(0.0, baseSun*dustFactor)

	baseUV := BaseUV + UVAmp*math.Cos((ls-250.0)*math.Pi/180.0)
	p.UVRadiation = math.Max(0.0, baseUV*dustFactor)

	// Dust and wind
	p.updateWindAndDust(dayFraction)

	// Seismicity and storms
	p.updateSeismicity()
	p.manageDustStorms()
}

func (p *Planet) updateWindAndDust(dayFraction float64) {
	if p.IsDustStorm {
		p.Dust += p.randomVariation(50, 150) * p.StormIntensity
		p.Dust = math.Min(p.Dust, MaxDust) // Cap dust
		p.WindSpeed = p.randomVariation(25, 40) * p.StormIntensity
	} else {
		if dayFraction > 0.4 && dayFraction < 0.7 { // Afternoon dust devils
			p.Dust = BaseDust + DustDevilExtra + p.randomVariation(-10, 10)
			p.WindSpeed = p.randomVariation(10, 25)
		} else { // Nighttime or calmer periods
			p.Dust = BaseDust + p.randomVariation(-10, 10)
			p.WindSpeed = p.randomVariation(5, 15)
		}
	}
}

func (p *Planet) manageDustStorms() {
	if p.IsDustStorm {
		p.StormDuration--
		// Gradual decrease in storm intensity and dust levels
		p.StormIntensity = math.Max(0.5, p.StormIntensity-0.01)
		p.Dust -= p.randomVariation(30, 70) // Gradually reduce dust
		p.Dust = math.Max(BaseDust, p.Dust) // Ensure dust doesn't drop below base
		if p.StormDuration == 0 {
			p.IsDustStorm = false
		}
	} else {
		p.Dust = math.Max(BaseDust, p.Dust-p.randomVariation(10, 50))
		if p.rng.Float64() < StormChance {
			p.IsDustStorm = true
			p.StormDuration = uint(p.rng.Intn(30) + 20) // Shorter storms
			p.StormIntensity = p.randomVariation(0.8, 1.2)
		}
	}
}

func (p *Planet) updateSeismicity() {
	r := p.rng.Float64()
	switch {
	case r < StrongSeismicityProb:
		p.Seismicity = StrongSeismicity
	case r < MediumSeismicityProb:
		p.Seismicity = MediumSeismicity
	case r < WeakSeismicityProb:
		p.Seismicity = WeakSeismicity
	default:
		p.Seismicity = NoSeismicity
	}
}

func (p *Planet) randomVariation(min, max float64) float64 {
	if min == max {
		return min
	}
	return min + (max-min)*p.rng.Float64()
}

func (p *Planet) PrintStatus() {
	fmt.Printf("Sol:%d %.2f:%.2f Ls=%.1f° | Temp: %.1f°C, Wind: %.1f m/s, Press: %.3f hPa, Dust: %.1f µg/m³, Sun: %.1f W/m², UV: %.1f W/m², Seismicity: %s, Dust Storm: %t\n",
		p.Time.Sol, p.Time.Hour, p.Time.Min, p.getLs(),
		p.Temperature, p.WindSpeed, p.Pressure, p.Dust, p.SunRadiation, p.UVRadiation, p.Seismicity, p.IsDustStorm)
}