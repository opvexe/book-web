// CodeMirror, copyright (c) by Marijn Haverbeke and others
// Distributed under an MIT license: http://codemirror.net/LICENSE

(function() {
  var mode = CodeMirror.getMode({tabSize: 4, indentUnit: 2}, "haml");
  function MT(name) { test.mode(name, mode, Array.prototype.slice.call(arguments, 1)); }

  // Requires at least one media query
  MT("elementName",
     "[tag %h1] Hey There");

  MT("oneElementPerLine",
     "[tag %h1] Hey There %h2");

  MT("idSelector",
     "[tag %h1][attribute #test] Hey There");

  MT("classSelector",
     "[tag %h1][attribute .hello] Hey There");

  MT("docType",
     "[tag !!! XML]");

  MT("comment",
     "[common / Hello WORLD]");

  MT("notComment",
     "[tag %h1] This is not a / common ");

  MT("attributes",
     "[tag %a]([variable title][operator =][string \"test\"]){[atom :title] [operator =>] [string \"test\"]}");

  MT("htmlCode",
     "[tag&bracket <][tag h1][tag&bracket >]Title[tag&bracket </][tag h1][tag&bracket >]");

  MT("rubyBlock",
     "[operator =][variable-2 @item]");

  MT("selectorRubyBlock",
     "[tag %a.selector=] [variable-2 @item]");

  MT("nestedRubyBlock",
      "[tag %a]",
      "   [operator =][variable puts] [string \"test\"]");

  MT("multilinePlaintext",
      "[tag %p]",
      "  Hello,",
      "  World");

  MT("multilineRuby",
      "[tag %p]",
      "  [common -# this is a common]",
      "     [common and this is a common too]",
      "  Date/Time",
      "  [operator -] [variable now] [operator =] [tag DateTime][operator .][property now]",
      "  [tag %strong=] [variable now]",
      "  [operator -] [keyword if] [variable now] [operator >] [tag DateTime][operator .][property parse]([string \"December 31, 2006\"])",
      "     [operator =][string \"Happy\"]",
      "     [operator =][string \"Belated\"]",
      "     [operator =][string \"Birthday\"]");

  MT("multilineComment",
      "[common /]",
      "  [common Multiline]",
      "  [common Comment]");

  MT("hamlComment",
     "[common -# this is a common]");

  MT("multilineHamlComment",
     "[common -# this is a common]",
     "   [common and this is a common too]");

  MT("multilineHTMLComment",
    "[common <!--]",
    "  [common what a common]",
    "  [common -->]");

  MT("hamlAfterRubyTag",
    "[attribute .block]",
    "  [tag %strong=] [variable now]",
    "  [attribute .test]",
    "     [operator =][variable now]",
    "  [attribute .right]");

  MT("stretchedRuby",
     "[operator =] [variable puts] [string \"Hello\"],",
     "   [string \"World\"]");

  MT("interpolationInHashAttribute",
     //"[tag %div]{[atom :id] [operator =>] [string \"#{][variable test][string }_#{][variable ting][string }\"]} test");
     "[tag %div]{[atom :id] [operator =>] [string \"#{][variable test][string }_#{][variable ting][string }\"]} test");

  MT("interpolationInHTMLAttribute",
     "[tag %div]([variable title][operator =][string \"#{][variable test][string }_#{][variable ting]()[string }\"]) Test");
})();
