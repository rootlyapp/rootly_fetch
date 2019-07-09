package models

// Plant model
type Plant struct {
	ID               int         `json:"id"`
	CommonName       string      `json:"common_name"`
	CompleteData     bool        `json:"complete_data"`
	Order            Category    `json:"order"`
	ScientificName   string      `json:"scientific_name"`
	Images           []Images    `json:"images"`
	FamilyCommonName string      `json:"family_common_name"`
	Family           Family      `json:"family"`
	Duration         string      `json:"duration"`
	Division         Category    `json:"division"`
	Class            Category    `json:"class"`
	Genus            Category    `json:"genus"`
	MainSpecies      MainSpecies `json:"main_species"`
}

// Images model
type Images struct {
	URL string `json:"url"`
}

// Height model
type Height struct {
	Ft float64 `json:"ft"`
	Cm float64 `json:"cm"`
}

// Density model
type Density struct {
	Sqm  float64 `json:"sqm"`
	Acre float64 `json:"acre"`
}

// Temperature model
type Temperature struct {
	DegF float64 `json:"deg_f"`
	DegC float64 `json:"deg_c"`
}

// Measure model
type Measure struct {
	Inches float64 `json:"inches"`
	Cm     float64 `json:"cm"`
}

// Category model
type Category struct {
	Slug string `json:"slug"`
	Name string `json:"name"`
	Link string `json:"link"`
	ID   int    `json:"id"`
}

// Family model
type Family struct {
	Category
	CommonName string `json:"common_name"`
}

// Flower model
type Flower struct {
	Conspicuous bool   `json:"conspicuous"`
	Color       string `json:"color"`
}

// SoilsAdaptation model
type SoilsAdaptation struct {
	Medium bool `json:"medium"`
	Fine   bool `json:"fine"`
	Coarse bool `json:"coarse"`
}

// Foliage model
type Foliage struct {
	Texture        string `json:"texture"`
	PorosityWinter string `json:"porosity_winter"`
	PorositySummer string `json:"porosity_summer"`
	Color          string `json:"color"`
}

// Propagation model
type Propagation struct {
	Tubers    bool `json:"tubers"`
	Sprigs    bool `json:"sprigs"`
	Sod       bool `json:"sod"`
	Seed      bool `json:"seed"`
	Cuttings  bool `json:"cuttings"`
	Corms     bool `json:"corms"`
	Container bool `json:"container"`
	Bulbs     bool `json:"bulbs"`
	BareRoot  bool `json:"bare_root"`
}

// Seed model
type Seed struct {
	VegetativeSpreadRate   string  `json:"vegetative_spread_rate"`
	SmallGrain             bool    `json:"small_grain"`
	SeedsPerPound          float64 `json:"seeds_per_pound"`
	SeedlingVigor          string  `json:"seedling_vigor"`
	SeedSpreadRate         string  `json:"seed_spread_rate"`
	CommercialAvailability string  `json:"commercial_availability"`
	BloomPeriod            string  `json:"bloom_period"`
}

// FruitOrSeed model
type FruitOrSeed struct {
	SeedPersistence bool   `json:"seed_persistence"`
	SeedPeriodEnd   string `json:"seed_period_end"`
	SeedPeriodBegin string `json:"seed_period_begin"`
	SeedAbundance   string `json:"seed_abundance"`
	Conspicuous     bool   `json:"conspicuous"`
	Color           string `json:"color"`
}

// Products model
type Products struct {
	Veneer                bool   `json:"veneer"`
	Pulpwood              bool   `json:"pulpwood"`
	ProteinPotential      string `json:"protein_potential"`
	Post                  bool   `json:"post"`
	PalatableHuman        bool   `json:"palatable_human"`
	PalatableBrowseAnimal string `json:"palatable_browse_animal"`
	PalatableGrazeAnimal  string `json:"palatable_graze_animal"`
	NurseryStock          bool   `json:"nursery_stock"`
	NavalStore            bool   `json:"naval_store"`
	Lumber                bool   `json:"lumber"`
	Fuelwood              string `json:"fuelwood"`
	Fodder                bool   `json:"fodder"`
	ChristmasTree         bool   `json:"christmas_tree"`
	BerryNutSeed          bool   `json:"berry_nut_seed"`
}

// Growth model
type Growth struct {
	TemperatureMinimum         Temperature `json:"temperature_minimum"`
	ShadeTolerance             string      `json:"shade_tolerance"`
	SalinityTolerance          string      `json:"salinity_tolerance"`
	RootDepthMinimum           Measure     `json:"root_depth_minimum"`
	ResproutAbility            bool        `json:"resprout_ability"`
	PrecipitationMinimum       Measure     `json:"precipitation_minimum"`
	PrecipitationMaximum       Measure     `json:"precipitation_maximum"`
	PlantingDensityMinimum     Measure     `json:"planting_density_minimum"`
	PlantingDensityMaximum     Measure     `json:"planting_density_maximum"`
	PhMinimum                  float64     `json:"ph_minimum"`
	PhMaximum                  float64     `json:"ph_maximum"`
	MoistureUse                string      `json:"moisture_use"`
	HedgeTorelance             string      `json:"hedge_tolerance"`
	FireTorelance              string      `json:"fire_tolerance"`
	FertilityRequirement       string      `json:"fertility_requirement"`
	DroughtTolerance           string      `json:"drought_tolerance"`
	ColdStratificationRequired bool        `json:"cold_stratification_required"`
	Caco3Tolerance             string      `json:"caco_3_tolerance"`
	AnaerobicTolerance         string      `json:"anaerobic_tolerance"`
	FrostFreeDaysMinimum       float64     `json:"frost_free_days_minimum"`
}

// Source model
type Source struct {
	SpeciesID  int    `json:"species_id"`
	SourceURL  string `json:"source_url"`
	Name       string `json:"name"`
	LastUpdate string `json:"last_update"`
}

// Specification model
type Specification struct {
	Toxicity            string `json:"toxicity"`
	ShapeAndOrientation string `json:"shape_and_orientation"`
	RegrowthRate        string `json:"regrowth_rate"`
	NitrogenFixation    string `json:"nitrogen_fixation"`
	MaxHeightAtBaseAge  Height `json:"max_height_at_base_age"`
	MatureHeight        Height `json:"mature_height"`
	LowGrowingGrass     bool   `json:"low_growing_grass"`
	Lifespan            string `json:"lifespan"`
	LeafRetention       bool   `json:"leaf_retention"`
	KnownAllelopath     bool   `json:"known_allelopath"`
	GrowthRate          string `json:"growth_rate"`
	GrowthPeriod        string `json:"growth_period"`
	GrowthHabit         string `json:"growth_habit"`
	GrowthForm          string `json:"growth_form"`
	FireResistance      bool   `json:"fire_resistance"`
	FallConspicuous     bool   `json:"fall_conspicuous"`
	CoppicePotential    bool   `json:"coppice_potential"`
	CNRatio             string `json:"c_n_ratio"`
	Bloat               string `json:"bloat"`
}

// MainSpecies model
type MainSpecies struct {
	Year            int             `json:"year"`
	Type            string          `json:"type"`
	Synonym         bool            `json:"synonym"`
	Status          string          `json:"status"`
	Specifications  Specification   `json:"specifications"`
	Sources         []Source        `json:"sources"`
	SoilsAdaptation SoilsAdaptation `json:"soils_adaptation"`
	Slug            string          `json:"slug"`
	Seed            Seed            `json:"seed"`
	ScientificName  string          `json:"scientific_name"`
	Propagation     Propagation     `json:"propagation"`
	Products        Products        `json:"products"`
	IsMainSpecies   bool            `json:"is_main_species"`
	Images          []Images        `json:"images"`
	Growth          Growth          `json:"growth"`
	FruitOrSeed     FruitOrSeed     `json:"fruit_or_seed"`
	Foliage         Foliage         `json:"foliage"`
	Flower          Flower          `json:"flower"`
	CompleteData    bool            `json:"complete_data"`
	Biography       string          `json:"biography"`
	Author          string          `json:"author"`
}
