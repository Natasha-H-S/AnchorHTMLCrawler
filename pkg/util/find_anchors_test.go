package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindAnchors(t *testing.T) {
	var tt = []struct {
		name     string
		body     string
		expected []string
	}{
		{
			name: "1. Returns Singular Anchor with no other attributes",
			body: `
			<body>
				<a href="abc">To ABC</a>
			</body>`,
			expected: []string{"abc"},
		},
		{
			name: "2. Returns Singular Anchor that has another attribute in the anchor after href",
			body: `
			<body>
				<a href="def" target="_blank">To DEF</a>
			</body>`,
			expected: []string{"def"},
		},
		{
			name: "3. Returns Singular Anchor that has another attribute in the anchor before href",
			body: `
			<body>
				<a target="_blank" href="ghi">To GHI</a>
			</body>`,
			expected: []string{"ghi"},
		},
		{
			name: "4. Returns nil if anchor tag has no href but other attribute",
			body: `
			<body>
				<a target="_blank">To JKL</a>
			</body>`,
			expected: nil,
		},
		{
			name: "5. Returns Multiple Anchors with other elements.",
			body: `
			<body>
				<a href="mno">To MNO</a>
				<span>SPAN ELEMENT</span>
				<li>
					<a href="pqr">To PQR</a>
				</li>
				<a href="stu">To STU</a>
			</body>`,
			expected: []string{"mno", "pqr", "stu"},
		},
		{
			name: "6. No Anchors found in URL",
			body: `
			<body>
				<h1>No Anchors</h1>
			</body>`,
			expected: nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := FindAnchors(tc.body)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
