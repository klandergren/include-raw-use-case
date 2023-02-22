---
layout: post
title:  "Liquid Syntax Error"
date:   2023-02-22 10:52:23 -0800
---

Removing the `raw / endraw` from the following:

{% raw %}
{% highlight go %}{% include golang/main.go %}{%- endhighlight -%}
{% endraw %}

will error with:

{% raw %}
Liquid Exception: Liquid syntax error (/path/to/include-raw-use-case/_includes/golang/main.go line 30): Variable '{{3, 2}' was not properly terminated with regexp: /\}\}/ included in /path/to/include-raw-use-case/_posts/2023-02-22-liquid-syntax-error.markdown
{% endraw %}

See [Include Raw No Error]({% post_url 2023-02-22-include-raw-no-error %}) for
alternative.
