class Solution {
    /**
     * @param {number[]} prices
     * @return {number}
     */
    maxProfit(prices) {
        let result = 0;

        for (let i = 0; i < prices.length; i++) {
            for (let j = i + 1; j < prices.length; j++) {
                const profit = prices[j] - prices[i];
                result = Math.max(result, profit);
            }
        }

        if (result <= 0) {
            return 0
        }

        return result;
    }
}
