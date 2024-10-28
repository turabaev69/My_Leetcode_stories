class Solution {
    public int numSub(String s) {
        long count = 0, result = 0;
        int mod = 1_000_000_007;
        
        for (char c : s.toCharArray()) {
            if (c == '1') {
                count++;
                result = (result + count) % mod;
            } else {
                count = 0;
            }
        }
        
        return (int) result;
    }
}
