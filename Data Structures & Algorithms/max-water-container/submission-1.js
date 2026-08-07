class Solution {
    /**
     * @param {number[]} heights
     * @return {number}
     */
    maxArea(heights) {
        let l = 0;
        let r = heights.length - 1;
        let result = 0;

        while (l < r) {
            const lHeight = heights[l];
            const rHeight = heights[r];
            const area = (r - l) * Math.min(lHeight, rHeight);
            result = Math.max(result, area);

            if (lHeight < rHeight) {
                l++
            } else {
                r--
            }
        }

        return result;
    }

}
