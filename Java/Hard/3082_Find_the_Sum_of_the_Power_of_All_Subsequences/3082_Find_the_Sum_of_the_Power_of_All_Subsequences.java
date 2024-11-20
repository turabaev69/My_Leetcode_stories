class Solution {
    private static final int MOD = 1_000_000_007;

    public int sumOfPower(int[] nums, int k) {
        int n = nums.length;
        long result = 0;

        for (int mask = 1; mask < (1 << n); mask++) {
            long sum = 0;
            int count = 0;

            for (int i = 0; i < n; i++) {
                if ((mask & (1 << i)) != 0) {
                    sum += nums[i];
                    count++;
                }
            }

            if (sum == k) {
                long power = (1L << count) - 1; // Total subsequences minus empty subset
                result = (result + power) % MOD;
            }
        }

        return (int) result;
    }
}
