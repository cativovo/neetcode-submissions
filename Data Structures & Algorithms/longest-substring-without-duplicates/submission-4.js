class Solution {
    /**
     * @param {string} s
     * @return {number}
     */
    lengthOfLongestSubstring(s) {
        if (s.length === 0) {
            return 0
        }

        let l = 0;
        let r = 1;
        let result = 1;

        while (r < s.length) {
            if (s.substring(l, r).includes(s[r])) {
                result = Math.max(result, r-l)
                l++
                r = l+1
                continue
            }

            r++
            result = Math.max(result, r-l)
        }

        return result
    }
}
