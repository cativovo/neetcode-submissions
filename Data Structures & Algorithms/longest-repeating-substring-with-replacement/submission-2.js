class Solution {
    /**
     * @param {string} s
     * @param {number} k
     * @return {number}
     */
    characterReplacement(s, k) {
        let result = 0
        let charSet = new Set(s)

        for (let c of charSet) {
            let count = 0;
            let l = 0;

            for (let r = 0; r < s.length; r++) {
                if (s[r] === c) {
                    count++
                }

                while ((r - l + 1) - count > k) {
                    if (s[l] === c) {
                        count--;
                    }
                    l++
                }

                result = Math.max(result, r - l + 1)
            }

        }

        return result
    }
}
