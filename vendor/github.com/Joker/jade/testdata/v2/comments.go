// Code generated by "jade.go"; DO NOT EDIT.

package jade

import (
	pool "github.com/valyala/bytebufferpool"
)

func tpl_comments(buffer *pool.ByteBuffer) {

	buffer.WriteString(`<!--  just some paragraphs --><p>foo</p><p>bar</p><p>foo</p><p>bar</p><body><!-- 
    Comments for your HTML readers.
    Use as much text as you want.

 --></body><!DOCTYPE html><!--[if IE 8]><html lang="en" class="lt-ie9"><![endif]--><!--[if gt IE 8]><!--><html lang="en"><!--<![endif]--><body><p>Supporting old web browsers is a pain.</p></body></html>`)

}
