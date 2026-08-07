class Solution {
    /**
     * @param {number[]} prices
     * @return {number}
     */
    maxProfit(prices) {
               let l = 0;
        let r = 1;
        let result = 0;

        while (r < prices.length) {
            if (prices[l] > prices[r]) {
                l = r;
                r++;
                continue;
            }

            const profit = prices[r] - prices[l];
            result = Math.max(result, profit);
            r++;
        }

        return result;

    }
}
