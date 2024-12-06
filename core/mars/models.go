package mars

import (
	"math/rand"
)

type SeismicityType string

const (
	NoSeismicity     SeismicityType = "no"
	WeakSeismicity   SeismicityType = "weak"
	MediumSeismicity SeismicityType = "medium"
	StrongSeismicity SeismicityType = "strong"
)

const (
	StormChance         = 0.02  // ~2% chance per update
	StrongSeismicityProb = 0.0005
	MediumSeismicityProb = 0.002
	WeakSeismicityProb   = 0.01
)

type Planet struct {
	Temperature    float64       // °C
	WindSpeed      float64       // m/s
	Pressure       float64       // hPa
	Dust           float64       // µg/m³
	SunRadiation   float64       // W/m²
	UVRadiation    float64       // W/m²
	Seismicity     SeismicityType
	IsDustStorm    bool          // Indicates if a dust storm is active
	StormDuration  uint          // Sols remaining in the current storm
	StormIntensity float64       // Intensity of the storm (affects dust and wind)
	Time           MartianTime
	rng            *rand.Rand // Random number generator
}

type MartianTime struct {
	Sol  uint    // Count of sols since start of simulation
	Hour float64 // 0 to ~24.659722 hours per sol
	Min  float64 // 0 to 59.9999 minutes
}

const (
	SolsInYear      = 668
	HoursInSol      = 24.659722
	MinsInHour      = 60.0
	BaseTemp        = -63.0
	TempAmplitude   = 20.0
	BasePressure    = 0.6
	PressureAmp     = 0.05
	BaseSun         = 590.0
	SunAmp          = 130.0
	BaseUV          = 10.0
	UVAmp           = 5.0
	BaseDust        = 100.0
	DustDevilExtra  = 50.0
	DustHighExtra   = 200.0
	MaxDust         = 10000.0 // µg/m³
)