class Solution {
    /**
     * @param {string} s
     * @param {string} t
     * @return {boolean}
     */
    isAnagram(s, t) {
            if (s.length !== t.length) {
      return false;
    }

    const sM = {};
    const tM = {};

    for (let i = 0; i < s.length; i++) {
      const sC = s[i];
      if (sM[sC]) {
        sM[sC] += 1;
      } else {
        sM[sC] = 1;
      }

      const tC = t[i];
      if (tM[tC]) {
        tM[tC] += 1;
      } else {
        tM[tC] = 1;
      }
    }

    for (const key in sM) {
      if (sM[key] !== tM[key]) {
        return false;
      }
    }

    return true;

    }
}
