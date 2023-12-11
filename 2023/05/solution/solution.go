package solution

import (
	"slices"
	"sort"
	"strconv"
	"strings"
	"sync"
)

type Mapping struct {
	SourceStart int
	SourceEnd   int
	Offset      int
}

type SourceLocation struct {
	LineStart int
	LineEnd   int
}

func LowestLocation(input []string, sourceLocs []SourceLocation) int {
	resultChan := make(chan int)
	var wg sync.WaitGroup

	seeds := parseSeeds(input[0])

	seedSoilMappings := ParseMappings(input[sourceLocs[0].LineStart:sourceLocs[0].LineEnd])
	soilFertilizerMappings := ParseMappings(input[sourceLocs[1].LineStart:sourceLocs[1].LineEnd])
	fertilizerWaterMappings := ParseMappings(input[sourceLocs[2].LineStart:sourceLocs[2].LineEnd])
	waterLightMappings := ParseMappings(input[sourceLocs[3].LineStart:sourceLocs[3].LineEnd])
	lightTemperatureMappings := ParseMappings(input[sourceLocs[4].LineStart:sourceLocs[4].LineEnd])
	temperatureHumidityMappings := ParseMappings(input[sourceLocs[5].LineStart:sourceLocs[5].LineEnd])
	humidityLocationMappings := ParseMappings(input[sourceLocs[6].LineStart:sourceLocs[6].LineEnd])

	for _, seed := range seeds {
		wg.Add(1)

		go func(seed int) {
			defer wg.Done()

			soil := FindDestination(seed, seedSoilMappings)
			fertilizer := FindDestination(soil, soilFertilizerMappings)
			water := FindDestination(fertilizer, fertilizerWaterMappings)
			light := FindDestination(water, waterLightMappings)
			temperate := FindDestination(light, lightTemperatureMappings)
			humidity := FindDestination(temperate, temperatureHumidityMappings)
			resultChan <- FindDestination(humidity, humidityLocationMappings)
		}(seed)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	results := []int{}
	for s := range resultChan {
		results = append(results, s)
	}

	return slices.Min(results)
}

func parseSeeds(s string) (seeds []int) {
	for _, field := range strings.Fields(strings.Split(s, ":")[1]) {
		val, _ := strconv.Atoi(field)
		seeds = append(seeds, val)
	}
	return seeds
}

func ParseMappings(input []string) []Mapping {
	mappings := make([]Mapping, len(input))

	for i, s := range input {
		fields := strings.Fields(s)

		sourceStart, _ := strconv.Atoi(fields[1])
		destintationStart, _ := strconv.Atoi(fields[0])
		rangeNum, _ := strconv.Atoi(fields[2])

		mappings[i] = Mapping{
			SourceStart: sourceStart,
			SourceEnd:   sourceStart + rangeNum - 1,
			Offset:      destintationStart - sourceStart,
		}
	}

	sort.Slice(mappings, func(i, j int) bool {
		return mappings[i].SourceStart < mappings[j].SourceStart
	})
	return mappings
}

func FindDestination(source int, mappings []Mapping) int {
	for _, mapping := range mappings {
		if source >= mapping.SourceStart && source <= mapping.SourceEnd {
			return source + mapping.Offset
		}
	}
	return source
}
