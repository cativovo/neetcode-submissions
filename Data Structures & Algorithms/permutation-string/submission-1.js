class Solution {
    /**
     * @param {string} s1
     * @param {string} s2
     * @return {boolean}
     */
    checkInclusion(s1, s2) {
        s1 = s1.split('').sort().join('')
        let l = 0;
        let r = s1.length

        while (r <= s2.length) {
           const f = s2.substring(l, r).split('').sort().join('')
           if (f === s1) {
            return true 
           }

           l++
           r++
        }

        return false
    }
}
