---
layout: post
title:  "Include Raw No Error"
date:   2023-02-22 10:52:24 -0800
---

Using the `include_raw` tag allows a file—who we may not be able / willing to
easily modify with `raw / endraw` tags (e.g. contained with a git submodule)—to
be included and highlighted without liquid parse errors.

{% highlight go %}{% include_raw golang/main.go %}{%- endhighlight -%}

See [Liquid Syntax Error]({% post_url 2023-02-22-liquid-syntax-error %}) for
existing behavior.
