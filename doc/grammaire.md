# Grammaire du MarkDown 

Ce document décrit la grammaire du GitHub Flavored Markdown.

```bnf

<titre> := "#" | "##" | "###" | "####" | "#####" | "######" <text>

<text>  := <emphasis> | <links> 

<emphasis> := <bold> | <italic>

<bold> := "**" <raw_text> "**"

<italic> := "*" <raw_text> "*"

```