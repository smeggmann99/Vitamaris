package base

type GenderType string

const (
	MaleGender   GenderType = "male" // Consumes more food and water, but is stronger
	FemaleGender GenderType = "female" // Consumes less food and water
)

func (gt GenderType) String() string {
	return string(gt)
}

type WorkerSpecialtyType string

const (
	EngineerSpecialty WorkerSpecialtyType = "engineer" // Takes care of the base's infrastructure
	BotanicSpecialty  WorkerSpecialtyType = "botanic"  // Takes care of the greenhouse, plants, and food production
	MedicSpecialty    WorkerSpecialtyType = "medic"    // Takes care of the crew's health, physical and mental, humans sometimes get ill
)

func (st WorkerSpecialtyType) String() string {
	return string(st)
}

type CleanlinessType string

const (
	HighCleanliness CleanlinessType = "high"
	MediumCleanliness CleanlinessType = "medium"
	LowCleanliness CleanlinessType = "low"
)

func (ct CleanlinessType) String() string {
	return string(ct)
}

type RoomType string

const (
	FiltersRoomType      RoomType = "filters"
	ReactorRoomType      RoomType = "reactor"
	PanicRoomType        RoomType = "panic"
	GreenhouseRoomType   RoomType = "greenhouse"
	BarracksRoomType     RoomType = "barracks"
	StorageRoomType      RoomType = "storage"
)

func (rt RoomType) String() string {
	return string(rt)
}

type Room struct {
	Type      RoomType
	Integrity float32 // 0-100%, the room's structural integrity, rooms break if it reaches 0%
}

type Filters struct {
	Room // gets worn over time, the more worn it is, the less oxygen and water it produces, needs maintenance engineers
	OxygenProduction  uint // liters per hour, the amount of oxygen the filters produce per hour
	WaterProduction   uint // liters per hour, the amount of water the filters produce per hour
	EnergyConsumption uint // kWh, the amount of energy the filters consume per hour
}

type Storage struct {
	Room // gets worn over time, the more worn it is, the less food, water, and energy it can hold, needs maintenance engineers
	Calories         uint // kcal, the amount of food the storage currently holds
	Water            uint // liters, the amount of water the storage currently holds
	Energy           uint // kWh, the amount of energy the storage currently holds
	CaloriesCapacity uint // kcal, the max amount of food in the closed loop system
	WaterCapacity    uint // liters, the max amount of water in the closed loop system
	EnergyCapacity   uint // kWh, the amount of energy the storage can hold (battery capacity)
}

type Barracks struct {
	Room // gets worn over time, the more worn it is, the less health it restores, needs maintenance medics
	WorkersCapacity uint // the number of workers the barracks can accommodate
	HealthPerHour   uint // %, the amount of health the barracks restores per hour per worker
}

type PanicRoom struct {
	Room
	WorkersCapacity uint // the number of workers the panic room can accommodate
}

type Greenhouse struct {
	Room // gets worn over time, the more worn it is, the less food it produces, needs maintenance botanics
	FoodProduction    uint // kcal per hour, the amount of food the greenhouse produces per hour
	EnergyConsumption uint // kWh, the amount of energy the greenhouse consumes per hour
}

type Solars struct {
	Room
	// kWh per hour, the amount of energy the solar panels produce per hour, depends on the sun's intensity and the panels' cleanliness
	EnergyProduction uint
	// affects the energy production, the dirtier the panels, the less energy they produce, affected by dust storms and the crew's work
	Cleanliness      CleanlinessType
}

type Base struct {
	Rooms       map[RoomType]Room // rooms in the base, one of each type
	Oxygen      uint // liters, affects the crew's health, humans need oxygen to breathe
	Humidity    uint // %, affects the crew's health, humans need a certain level of humidity to be comfortable
	Temperature uint // Celsius, affects the crew's health, humans need a certain temperature to be comfortable
}

type Worker struct {
	Name                string
	Health              uint // 0-100%, humans die at 0%
	Gender              GenderType // affects the crew's oxygen, food and water consumption
	Speciality		    WorkerSpecialtyType // affects the crew's work efficiency in case of different tasks
	OxygenConsumption   uint // liters per hour, the amount of oxygen a worker consumes per hour
	WaterConsumption    uint // liters per hour, the amount of water a worker consumes per hour
	CaloriesConsumption uint // kcal per hour, the amount of calories a worker consumes per hour
	InBaracks           bool // if the worker is in the barracks, they restore health per hour
}