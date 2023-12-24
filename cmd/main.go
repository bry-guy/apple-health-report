package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

// Grepped types from export.xml
// HKCategoryTypeIdentifierAppleStandHour
// HKCategoryTypeIdentifierAudioExposureEvent
// HKCategoryTypeIdentifierHeadphoneAudioExposureEvent
// HKCategoryTypeIdentifierHighHeartRateEvent
// HKCategoryTypeIdentifierSleepAnalysis
// HKDataTypeSleepDurationGoal
// HKQuantityTypeIdentifierActiveEnergyBurned
// HKQuantityTypeIdentifierAppleExerciseTime
// HKQuantityTypeIdentifierAppleStandTime
// HKQuantityTypeIdentifierAppleWalkingSteadiness
// HKQuantityTypeIdentifierBasalEnergyBurned
// HKQuantityTypeIdentifierBodyFatPercentage
// HKQuantityTypeIdentifierBodyMass
// HKQuantityTypeIdentifierBodyMassIndex
// HKQuantityTypeIdentifierDistanceCycling
// HKQuantityTypeIdentifierDistanceDownhillSnowSports
// HKQuantityTypeIdentifierDistanceWalkingRunning
// HKQuantityTypeIdentifierEnvironmentalAudioExposure
// HKQuantityTypeIdentifierEnvironmentalSoundReduction
// HKQuantityTypeIdentifierFlightsClimbed
// HKQuantityTypeIdentifierHeadphoneAudioExposure
// HKQuantityTypeIdentifierHeartRate
// HKQuantityTypeIdentifierHeartRateRecoveryOneMinute
// HKQuantityTypeIdentifierHeartRateVariabilitySDNN
// HKQuantityTypeIdentifierHeight
// HKQuantityTypeIdentifierLeanBodyMass
// HKQuantityTypeIdentifierOxygenSaturation
// HKQuantityTypeIdentifierPhysicalEffort
// HKQuantityTypeIdentifierRespiratoryRate
// HKQuantityTypeIdentifierRestingHeartRate
// HKQuantityTypeIdentifierRunningGroundContactTime
// HKQuantityTypeIdentifierRunningPower
// HKQuantityTypeIdentifierRunningSpeed
// HKQuantityTypeIdentifierRunningStrideLength
// HKQuantityTypeIdentifierRunningVerticalOscillation
// HKQuantityTypeIdentifierSixMinuteWalkTestDistance
// HKQuantityTypeIdentifierStairAscentSpeed
// HKQuantityTypeIdentifierStairDescentSpeed
// HKQuantityTypeIdentifierStepCount
// HKQuantityTypeIdentifierTimeInDaylight
// HKQuantityTypeIdentifierVO2Max
// HKQuantityTypeIdentifierWalkingAsymmetryPercentage
// HKQuantityTypeIdentifierWalkingDoubleSupportPercentage
// HKQuantityTypeIdentifierWalkingHeartRateAverage
// HKQuantityTypeIdentifierWalkingSpeed
// HKQuantityTypeIdentifierWalkingStepLength
// HKQuantityTypeIdentifierOxygenSaturation
// HKQuantityTypeIdentifierVO2Max
// HKQuantityTypeIdentifierDistanceDownhillSnowSports
// HKQuantityTypeIdentifierWalkingAsymmetryPercentage
// HKQuantityTypeIdentifierEnvironmentalSoundReduction
// HKQuantityTypeIdentifierEnvironmentalAudioExposure
// HKQuantityTypeIdentifierSixMinuteWalkTestDistance
// HKQuantityTypeIdentifierRunningVerticalOscillation
// HKQuantityTypeIdentifierHeartRateRecoveryOneMinute
// HKQuantityTypeIdentifierRunningPower
// HKQuantityTypeIdentifierTimeInDaylight
// HKCategoryTypeIdentifierAudioExposureEvent
// HKCategoryTypeIdentifierHeadphoneAudioExposureEvent
// HKCategoryTypeIdentifierHighHeartRateEvent

type Record struct {
	Type                         string                             `xml:"type,attr"`
	SourceName                   string                             `xml:"sourceName,attr"`
	SourceVersion                string                             `xml:"sourceVersion,attr"`
	Device                       string                             `xml:"device,attr"`
	Unit                         string                             `xml:"unit,attr"`
	CreationDate                 string                             `xml:"creationDate,attr"`
	StartDate                    string                             `xml:"startDate,attr"`
	EndDate                      string                             `xml:"endDate,attr"`
	Value                        string                             `xml:"value,attr"`
	MetadataEntries              []MetadataEntry                    `xml:"MetadataEntry"`
	HeartRateVariabilityMetadata []HeartRateVariabilityMetadataList `xml:"HeartRateVariabilityMetadataList"`
}

type MetadataEntry struct {
	Key   string `xml:"key,attr"`
	Value string `xml:"value,attr"`
}

type HeartRateVariabilityMetadataList struct {
	InstantaneousBPM []InstantaneousBeatsPerMinute `xml:"InstantaneousBeatsPerMinute"`
}
type InstantaneousBeatsPerMinute struct {
	BPM  int    `xml:"bpm,attr"`
	Time string `xml:"time,attr"`
}

type Me struct {
	DateOfBirth                 string `xml:"HKCharacteristicTypeIdentifierDateOfBirth,attr"`
	BiologicalSex               string `xml:"HKCharacteristicTypeIdentifierBiologicalSex,attr"`
	BloodType                   string `xml:"HKCharacteristicTypeIdentifierBloodType,attr"`
	FitzpatrickSkinType         string `xml:"HKCharacteristicTypeIdentifierFitzpatrickSkinType,attr"`
	CardioFitnessMedicationsUse string `xml:"HKCharacteristicTypeIdentifierCardioFitnessMedicationsUse,attr"`
}

type Correlation struct {
	Type            string          `xml:"type,attr"`
	SourceName      string          `xml:"sourceName,attr"`
	SourceVersion   string          `xml:"sourceVersion,attr"`
	Device          string          `xml:"device,attr"`
	CreationDate    string          `xml:"creationDate,attr"`
	StartDate       string          `xml:"startDate,attr"`
	EndDate         string          `xml:"endDate,attr"`
	MetadataEntries []MetadataEntry `xml:"MetadataEntry"`
	Records         []Record        `xml:"Record"`
}

type HealthData struct {
	ExportDate          ExportDate           `xml:"ExportDate"`
	Me                  Me                   `xml:"Me"`
	Records             []Record             `xml:"Record"`
	Correlations        []Correlation        `xml:"Correlation"`
	Workouts            []Workout            `xml:"Workout"`
	ActivitySummaries   []ActivitySummary    `xml:"ActivitySummary"`
	ClinicalRecords     []ClinicalRecord     `xml:"ClinicalRecord"`
	Audiograms          []Audiogram          `xml:"Audiogram"`
	VisionPrescriptions []VisionPrescription `xml:"VisionPrescription"`
}

type ExportDate struct {
	Value string `xml:"value,attr"`
}

type ClinicalRecord struct {
	Type             string `xml:"type,attr"`
	Identifier       string `xml:"identifier,attr"`
	SourceName       string `xml:"sourceName,attr"`
	SourceURL        string `xml:"sourceURL,attr"`
	FhirVersion      string `xml:"fhirVersion,attr"`
	ReceivedDate     string `xml:"receivedDate,attr"`
	ResourceFilePath string `xml:"resourceFilePath,attr"`
}

type Audiogram struct {
	Type              string             `xml:"type,attr"`
	SourceName        string             `xml:"sourceName,attr"`
	SourceVersion     string             `xml:"sourceVersion,attr"`
	Device            string             `xml:"device,attr"`
	CreationDate      string             `xml:"creationDate,attr"`
	StartDate         string             `xml:"startDate,attr"`
	EndDate           string             `xml:"endDate,attr"`
	SensitivityPoints []SensitivityPoint `xml:"SensitivityPoint"`
	MetadataEntries   []MetadataEntry    `xml:"MetadataEntry"`
}

type SensitivityPoint struct {
	FrequencyValue string `xml:"frequencyValue,attr"`
	FrequencyUnit  string `xml:"frequencyUnit,attr"`
	LeftEarValue   string `xml:"leftEarValue,attr"`
	LeftEarUnit    string `xml:"leftEarUnit,attr"`
	RightEarValue  string `xml:"rightEarValue,attr"`
	RightEarUnit   string `xml:"rightEarUnit,attr"`
}

type VisionPrescription struct {
	Type            string          `xml:"type,attr"`
	DateIssued      string          `xml:"dateIssued,attr"`
	ExpirationDate  string          `xml:"expirationDate,attr"`
	Brand           string          `xml:"brand,attr"`
	RightEye        RightEye        `xml:"RightEye"`
	LeftEye         LeftEye         `xml:"LeftEye"`
	Attachments     []Attachment    `xml:"Attachment"`
	MetadataEntries []MetadataEntry `xml:"MetadataEntry"`
}

type Eye struct {
	Sphere          string `xml:"sphere,attr"`
	SphereUnit      string `xml:"sphereUnit,attr"`
	Cylinder        string `xml:"cylinder,attr"`
	CylinderUnit    string `xml:"cylinderUnit,attr"`
	Axis            string `xml:"axis,attr"`
	AxisUnit        string `xml:"axisUnit,attr"`
	Add             string `xml:"add,attr"`
	AddUnit         string `xml:"addUnit,attr"`
	Vertex          string `xml:"vertex,attr"`
	VertexUnit      string `xml:"vertexUnit,attr"`
	PrismAmount     string `xml:"prismAmount,attr"`
	PrismAmountUnit string `xml:"prismAmountUnit,attr"`
	PrismAngle      string `xml:"prismAngle,attr"`
	PrismAngleUnit  string `xml:"prismAngleUnit,attr"`
	FarPD           string `xml:"farPD,attr"`
	FarPDUnit       string `xml:"farPDUnit,attr"`
	NearPD          string `xml:"nearPD,attr"`
	NearPDUnit      string `xml:"nearPDUnit,attr"`
	BaseCurve       string `xml:"baseCurve,attr"`
	BaseCurveUnit   string `xml:"baseCurveUnit,attr"`
	Diameter        string `xml:"diameter,attr"`
	DiameterUnit    string `xml:"diameterUnit,attr"`
}

type RightEye Eye
type LeftEye Eye

type Workout struct {
	WorkoutActivityType   string            `xml:"workoutActivityType,attr"`
	Duration              string            `xml:"duration,attr"`
	DurationUnit          string            `xml:"durationUnit,attr"`
	TotalDistance         string            `xml:"totalDistance,attr"`
	TotalDistanceUnit     string            `xml:"totalDistanceUnit,attr"`
	TotalEnergyBurned     string            `xml:"totalEnergyBurned,attr"`
	TotalEnergyBurnedUnit string            `xml:"totalEnergyBurnedUnit,attr"`
	SourceName            string            `xml:"sourceName,attr"`
	SourceVersion         string            `xml:"sourceVersion,attr"`
	Device                string            `xml:"device,attr"`
	CreationDate          string            `xml:"creationDate,attr"`
	StartDate             string            `xml:"startDate,attr"`
	EndDate               string            `xml:"endDate,attr"`
	MetadataEntries       []MetadataEntry   `xml:"MetadataEntry"`
	WorkoutEvents         []WorkoutEvent    `xml:"WorkoutEvent"`
	WorkoutStatistics     WorkoutStatistics `xml:"WorkoutStatistics"`
}

type WorkoutEvent struct {
	Type         string `xml:"type,attr"`
	Date         string `xml:"date,attr"`
	Duration     string `xml:"duration,attr"`
	DurationUnit string `xml:"durationUnit,attr"`
	// Additional attributes as needed
}

type WorkoutStatistics struct {
	Type      string `xml:"type,attr"`
	StartDate string `xml:"startDate,attr"`
	EndDate   string `xml:"endDate,attr"`
	Average   string `xml:"average,attr"`
	Minimum   string `xml:"minimum,attr"`
	Maximum   string `xml:"maximum,attr"`
	Sum       string `xml:"sum,attr"`
	Unit      string `xml:"unit,attr"`
}

type ActivitySummary struct {
	DateComponents         string `xml:"dateComponents,attr"`
	ActiveEnergyBurned     string `xml:"activeEnergyBurned,attr"`
	ActiveEnergyBurnedGoal string `xml:"activeEnergyBurnedGoal,attr"`
	ActiveEnergyBurnedUnit string `xml:"activeEnergyBurnedUnit,attr"`
	AppleMoveTime          string `xml:"appleMoveTime,attr"`
	AppleMoveTimeGoal      string `xml:"appleMoveTimeGoal,attr"`
	AppleExerciseTime      string `xml:"appleExerciseTime,attr"`
	AppleExerciseTimeGoal  string `xml:"appleExerciseTimeGoal,attr"`
	AppleStandHours        string `xml:"appleStandHours,attr"`
	AppleStandHoursGoal    string `xml:"appleStandHoursGoal,attr"`
}

type Attachment struct {
	Identifier string `xml:"identifier,attr"`
}

const (
	AppleStandHourStood = "HKCategoryValueAppleStandHourStood"
	AppleStandHourIdle  = "HKCategoryValueAppleStandHourIdle"
)

func main() {
	// Read the XML file
	fmt.Println("reading file")
	xmlData, err := os.ReadFile("exports/2023/bryan/export-headless.xml")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Parse the XML data
	fmt.Println("parsing file")
	var healthData HealthData
	err = xml.Unmarshal(xmlData, &healthData)
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return
	}

	moveDays := map[int]map[int]float64{}
	exerciseDays := map[int]map[int]float64{}
	standDays := map[int]map[int]int{}

	typesMap := map[string]int{}

	for _, record := range healthData.Records {
		layout := "2006-01-02 15:04:05 -0700"
		start, err := time.Parse(layout, record.StartDate)

		if err != nil {
			fmt.Printf("unable to time.Parse record.StartDate: %+v\n", record.StartDate)
			continue
		}

		if start.Year() != 2023 {
			continue
		}

		month := int(start.Month())
		day := start.Day()

		if moveDays[month] == nil {
			moveDays[month] = make(map[int]float64)
		}

		if exerciseDays[month] == nil {
			exerciseDays[month] = make(map[int]float64)
		}

		if standDays[month] == nil {
			standDays[month] = make(map[int]int)
		}

		// Sum values for Move, Exercise, and Stand Rings
		switch record.Type {
		case "HKQuantityTypeIdentifierActiveEnergyBurned":
			moveDays[month][day] += convertToFloat(record.Value)
		case "HKQuantityTypeIdentifierAppleExerciseTime":
			exerciseDays[month][day] += convertToFloat(record.Value)
		case "HKCategoryTypeIdentifierAppleStandHour":
			switch record.Value {
			case AppleStandHourIdle:
			case AppleStandHourStood:
				standDays[month][day] += 1
			}
		}

		typesMap[record.Type] += 1
	}

	moveGoal := 480
	exerciseGoal := 30
	standGoal := 12

	moveRingClosedCount := calculateClosedRings[float64]("move", moveDays, moveGoal)
	exerciseRingClosedCount := calculateClosedRings[float64]("exercise", exerciseDays, exerciseGoal)
	standRingClosedCount := calculateClosedRings[int]("stand", standDays, standGoal)

	// for t, c := range typesMap {
	// 	fmt.Printf("Type: %s, Count: %d\n", t, c)
	// }

	fmt.Printf("Move Ring Closed: %d\n", moveRingClosedCount)
	fmt.Printf("Exercise Ring Closed: %d\n", exerciseRingClosedCount)
	fmt.Printf("Stand Ring Closed: %d\n", standRingClosedCount)
	fmt.Printf("Total Rings Closed: %d\n", moveRingClosedCount+exerciseRingClosedCount+standRingClosedCount)
}

func convertToFloat(s string) float64 {
	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatal(err)
	}
	return val
}

type Sum interface {
	int | float64
}

func calculateClosedRings[T Sum](title string, categoryDays map[int]map[int]T, goal int) int {
	var closed int

	closedCount := 0

	for month, dayMap := range categoryDays {
		fmt.Printf("month: %d\n", month)
		for day, sum := range dayMap {
			closed = int(sum) / goal
			closedCount += int(sum) / goal
			fmt.Printf("day: %d\t%sSum: %v\tclosed: %d\n", day, title, sum, closed)
		}
	}
	return closedCount
}
