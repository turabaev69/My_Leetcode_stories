
func alertNames(keyName []string, keyTime []string) []string {
	usageMap := make(map[string][]int)
	for i := 0; i < len(keyName); i++ {
		timeParts := strings.Split(keyTime[i], ":")
		hours, _ := strconv.Atoi(timeParts[0])
		minutes, _ := strconv.Atoi(timeParts[1])
		totalMinutes := hours*60 + minutes
		usageMap[keyName[i]] = append(usageMap[keyName[i]], totalMinutes)
	}

	result := []string{}
	for name, times := range usageMap {
		sort.Ints(times)
		for i := 2; i < len(times); i++ {
			if times[i]-times[i-2] <= 60 {
				result = append(result, name)
				break
			}
		}
	}

	sort.Strings(result)
	return result
}
