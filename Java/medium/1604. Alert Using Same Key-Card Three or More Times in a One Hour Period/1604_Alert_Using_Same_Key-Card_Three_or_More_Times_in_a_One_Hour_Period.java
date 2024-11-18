import java.util.*;

class Solution {
    public List<String> alertNames(String[] keyName, String[] keyTime) {
        Map<String, List<Integer>> usageMap = new HashMap<>();
        for (int i = 0; i < keyName.length; i++) {
            usageMap.putIfAbsent(keyName[i], new ArrayList<>());
            String[] time = keyTime[i].split(":");
            int minutes = Integer.parseInt(time[0]) * 60 + Integer.parseInt(time[1]);
            usageMap.get(keyName[i]).add(minutes);
        }

        List<String> result = new ArrayList<>();
        for (String name : usageMap.keySet()) {
            List<Integer> times = usageMap.get(name);
            Collections.sort(times);
            for (int i = 2; i < times.size(); i++) {
                if (times.get(i) - times.get(i - 2) <= 60) {
                    result.add(name);
                    break;
                }
            }
        }

        Collections.sort(result);
        return result;
    }
}
