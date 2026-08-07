class Solution {
    /**
     * @param {string} s
     * @return {boolean}
     */
    isPalindrome(s) {
        s = s.toLowerCase()
        s = s.replaceAll(/[^a-z0-9]/g, "")
        let l = 0;
        let r = s.length - 1;
  
        while (l < r) {
            console.log(s[l], s[r])
            if (s[l] !== s[r]) {
                return false
            }

            l++
            r--
        }

        return true;
    }
}
