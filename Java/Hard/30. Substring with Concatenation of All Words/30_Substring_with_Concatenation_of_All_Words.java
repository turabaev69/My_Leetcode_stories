class Solution {
 public List<Integer> findSubstring(String s, String[] words) {
        List<Integer> result = new ArrayList<>();
        if (s == null || s.length() == 0 || words == null || words.length == 0) return result;

        int wordLength = words[0].length();
        int wordsCount = words.length;
        int substringLength = wordLength * wordsCount;

        Map<String, Integer> wordMap = new HashMap<>();
        for (String word : words) wordMap.put(word, wordMap.getOrDefault(word, 0) + 1);

        for (int i = 0; i < wordLength; i++) {
            int left = i;
            Map<String, Integer> seen = new HashMap<>();
            int count = 0;

            for (int j = i; j <= s.length() - wordLength; j += wordLength) {
                String word = s.substring(j, j + wordLength);

                if (wordMap.containsKey(word)) {
                    seen.put(word, seen.getOrDefault(word, 0) + 1);
                    count++;

                    while (seen.get(word) > wordMap.get(word)) {
                        String leftWord = s.substring(left, left + wordLength);
                        seen.put(leftWord, seen.get(leftWord) - 1);
                        count--;
                        left += wordLength;
                    }

                    if (count == wordsCount) result.add(left);
                } else {
                    seen.clear();
                    count = 0;
                    left = j + wordLength;
                }
            }
        }
        return result;
    }
}
