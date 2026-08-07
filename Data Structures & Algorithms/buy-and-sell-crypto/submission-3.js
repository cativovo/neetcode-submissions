class Solution {
    /**
     * @param {number[]} prices
     * @return {number}
     */
    maxProfit(prices) {
        let l = 0;
        let r = 1;

        if (prices.length < r) {
            return 0;
        }

        let result = 0;

        while (r < prices.length) {
            if (prices[l] > prices[r]) {
                l = r
                r++
                continue;
            }
            const profit = prices[r] - prices[l]; 
            result = Math.max(result, profit);
            r++

        }


        if (result <= 0) {
            return 0;
        }

        return result;
    }
}
