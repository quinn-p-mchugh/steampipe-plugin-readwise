package zendesk

import (
    "context"

    "github.com/turbot/steampipe-plugin-sdk/plugin"
    "github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
    p := &plugin.Plugin{
        Name:             "steampipe-plugin-readwise",
        DefaultTransform: transform.FromGo().NullIfZero(),
        TableMap: map[string]*plugin.Table{
            "readwise_book":        tableReadwiseBook(),
            "readwise_highlight":        tableReadwiseHighlight(),
            "readwise_highlight_tag": tableReadwiseHighlightTag(),
            "readwise_book_tag":       tableReadwiseBookTag(),
            "readwise_highlight_tag_map":       tableReadwiseHighlightTagMap(),
            "readwise_book_tag_map": tableReadwiseBookTagMap(),
        },
    }
    return p
}