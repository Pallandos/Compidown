# Grammaire du MarkDown 

Ce document décrit la grammaire du GitHub Flavored Markdown.

Pour référence, la documentation utilisée est lisible [ici](https://github.github.com/gfm/).

> Dans la grammaire ci dessous, je ne tient pas compte des espaces / retours à la ligne. Je ne les prendrais en compte que dans la phase d'analyse lexicale avec les REGEX, sinon cela ne ferait que rendre la grammaire peu digeste.

```ebnf

<document>      := { <block> }

<block>         := <leaf_block> | <cont_block> 

<leaf_block>    := <theme_break> 
                | <titre> 
                | <indent_code>
                | <fenced_code>
                | <html_block>
                | <paragraph>
                | <blank>
                | <table>

<cont_block>    := <quote>
                | <list>

<text>          := { <inline> }

(* containers blocks : *)

<quote>         := ">" <text>

<list>          := { <list_item> }

<list_item>     := <list_marker> <list> | <text>

<list_marker>   := <ordered_l> | <bullet>

<ordered_l>     := ("0" | "1" | "2" | "3" | "4" | "5" 
                | "6" | "7" | "8" | "9") ("." | ")")

<bullet>        := "+" | "-" | "*"

(* leaf blocks *)

<titre>         := "#" 
                | "##" 
                | "###" 
                | "####" 
                | "#####" 
                | "######" <text>

<indent_code>   := ?tabulation? <raw_text>

<fenced_code>   := "```" <raw_text> <raw_text> "```"

<paragraph>     := { <text> }

(* inlines *)

<inline>        := <raw_text>
                | <emphasis>
                | <link>
                | <image>

<emphasis>      := <bold> | <italic> | <code_emph>

<bold>          := "**" <raw_text> "**"

<italic>        := "*" <raw_text> "*"

<code_emph>     := "`" <raw_text> "`"

<link>          := "[" <text> "](" <url> ")"

<image>         := "![" <raw_text> "](" <url> ")"  

```
