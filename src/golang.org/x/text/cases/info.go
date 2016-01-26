// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cases

func (c info) cccVal() info {
	if c&exceptionBit != 0 {
		// TODO: return info(exceptions[c>>exceptionShift]) & cccMask
		return cccZero
	}
	return c & cccMask
}

func (c info) cccType() info {
	ccc := c.cccVal()
	if ccc <= cccZero {
		return cccZero
	}
	return ccc
}

// For the purpose of title casing we use an approximation of the Unicode Word
// Breaking algorithm defined in Annex #29:
// http://www.unicode.org/reports/tr29/#Default_Grapheme_Cluster_Table.
//
// For our approximation, we group the Word Break types into the following
// categories, with associated rules:
//
// 1) Letter:
//    ALetter, Hebrew_Letter, Numeric, ExtendNumLet, Extend.
//    Rule: Never break between consecutive runes of this category.
//
// 2) Mid:
//    Format, MidLetter, MidNumLet, Single_Quote.
//    (Cf. case-ignorable: MidLetter, MidNumLet or cat is Mn, Me, Cf, Lm or Sk).
//    Rule: Don't break between Letter and Mid, but break between two Mids.
//
// 3) Break:
//    Any other category, including NewLine, CR, LF and Double_Quote. These
//    categories should always result in a break between two cased letters.
//    Rule: Always break.
//
// Note 1: the Katakana and MidNum categories can, in esoteric cases, result in
// preventing a break between two cased letters. For now we will ignore this
// (e.g. [ALetter] [ExtendNumLet] [Katakana] [ExtendNumLet] [ALetter] and
// [ALetter] [Numeric] [MidNum] [Numeric] [ALetter].)
//
// Note 2: the rule for Mid is very approximate, but works in most cases. To
// improve, we could store the categories in the trie value and use a FA to
// manage breaks. See TODO comment above.
//
// Note 3: according to the spec, it is possible for the Extend category to
// introduce breaks between other categories grouped in Letter. However, this
// is undesirable for our purposes. ICU prevents breaks in such cases as well.

// isBreak returns whether this rune should introduce a break.
func (c info) isBreak() bool {
	return c.cccVal() == cccBreak
}

// isLetter returns whether the rune is of break type ALetter, Hebrew_Letter,
// Numeric, ExtendNumLet, or Extend.
func (c info) isLetter() bool {
	ccc := c.cccVal()
	if ccc == cccZero {
		return !c.isCaseIgnorable()
	}
	return ccc != cccBreak
}
