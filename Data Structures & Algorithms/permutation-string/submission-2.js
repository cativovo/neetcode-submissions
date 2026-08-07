class Solution {
    /**
     * @param {string} s1
     * @param {string} s2
     * @return {boolean}
     */
    checkInclusion(s1, s2) {
        const m1 = new Map()
        for (let i = 0; i < s1.length; i++) {
            const k = s1[i]
            const n = m1.get(k) ?? 0
            m1.set(k, n + 1)
        }

        let l = 0;
        let r = s1.length

        while (r <= s2.length) {
            const s = s2.substring(l, r)
            const m2 = new Map()
            for (let i = 0; i < s.length; i++) {
                const k = s[i]
                const n = m2.get(k) ?? 0
                m2.set(k, n + 1)
            }

            let valid = false

            for (const k of m1.keys()) {
               if (m1.get(k) != m2.get(k)) {
                valid = false
                break
               }
               valid = true
            }

            if (valid) {
                return true
            }

            l++
            r++
        }

        return false
    }
}
