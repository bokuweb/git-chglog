package chglog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGitHubProcessor(t *testing.T) {
	assert := assert.New(t)

	config := &Config{
		Info: &Info{
			RepositoryURL: "https://example.com",
		},
	}

	processor := &GitHubProcessor{}

	processor.Bootstrap(config)

	assert.Equal(
		&Commit{
			Header:  "message [@foo](https://github.com/foo) [#123](https://example.com/issues/123)",
			Subject: "message [@foo](https://github.com/foo) [#123](https://example.com/issues/123)",
			Body: `issue [#456](https://example.com/issues/456)
multiline [#789](https://example.com/issues/789)
[@foo](https://github.com/foo), [@bar](https://github.com/bar)`,
			Notes: []*Note{
				&Note{
					Body: `issue1 [#11](https://example.com/issues/11)
issue2 [#22](https://example.com/issues/22)
[gh-56](https://example.com/issues/56) hoge fuga`,
				},
			},
		},
		processor.ProcessCommit(
			&Commit{
				Header:  "message @foo #123",
				Subject: "message @foo #123",
				Body: `issue #456
multiline #789
@foo, @bar`,
				Notes: []*Note{
					&Note{
						Body: `issue1 #11
issue2 #22
gh-56 hoge fuga`,
					},
				},
			},
		),
	)
}
