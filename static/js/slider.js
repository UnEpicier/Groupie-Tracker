let $range = document.getElementByClassName("js-range-slider"),
    $from = document.getElementByClassName("from"),
    $to = document.getElementByClassName("to"),
    range,
    min = $range.data('min'),
    max = $range.data('max'),
    from,
    to;

let updateValues = function () {
    $from.prop("value", from);
    $to.prop("value", to);
};

$range.ionRangeSlider({
    onChange: function (data) {
        from = data.from;
        to = data.to;
        updateValues();
    }
});

range = $range.data("ionRangeSlider");
let updateRange = function () {
    range.update({
        from: from,
        to: to
    });
};

$from.on("input", function () {
    from = +this.prop("value");
    if (from < min) {
        from = min;
    }
    if (from > to) {
        from = to;
    }
    updateValues();
    updateRange();
});

$to.on("input", function () {
    to = +this.prop("value");
    if (to > max) {
        to = max;
    }
    if (to < from) {
        to = from;
    }
    updateValues();
    updateRange();
});