$(function() {

    class GaugeChart {
        constructor(element, params) {
            this._element = element;
            this._initialValue = params.initialValue;
            this._higherValue = params.higherValue;
            this._title = params.title;
            this._subtitle = params.subtitle;
        }

        _buildConfig() {
            let element = this._element;

            return {
                value: this._initialValue,
                valueIndicator: {
                    color: '#fff'
                },
                geometry: {
                    startAngle: 180,
                    endAngle: 360
                },
                scale: {
                    startValue: 0,
                    endValue: this._higherValue,
                    customTicks: [0, 200, 500, 800, 1000, 1500, 2000, 3000],
                    tick: {
                        length: 8
                    },
                    label: {
                        font: {
                            color: '#87959f',
                            size: 9,
                            family: '"Open Sans", sans-serif'
                        }
                    }
                },
                title: {
                    verticalAlignment: 'bottom',
                    text: this._title,
                    font: {
                        family: '"Open Sans", sans-serif',
                        color: '#fff',
                        size: 10
                    },
                    subtitle: {
                        text: this._subtitle,
                        font: {
                            family: '"Open Sans", sans-serif',
                            color: '#fff',
                            weight: 700,
                            size: 28
                        }
                    }
                },
                onInitialized: function() {
                    let currentGauge = $(element);
                    let circle = currentGauge.find('.dxg-spindle-hole').clone();
                    let border = currentGauge.find('.dxg-spindle-border').clone();

                    currentGauge.find('.dxg-title text').first().attr('y', 48);
                    currentGauge.find('.dxg-title text').last().attr('y', 28);
                    currentGauge.find('.dxg-value-indicator').append(border, circle);
                }

            }
        }

        init() {
            $(this._element).dxCircularGauge(this._buildConfig());
        }
    }

    function updateCo2Handler() {
        $.ajax({
            url: "http://sensors.sahnovsky.life/event/latest",
            headers: {"Content-Type": "application/json"},
            context: document.body
        }).done(function(data) {
            let item = $('#co2');
            let gauge = $(item).dxCircularGauge('instance');
            let value = data.cO2;
            let gaugeElement = $(gauge._$element[0]);

            gaugeElement.find('.dxg-title text').last().html(`${value}`);
            gauge.value(value);

            if (data.cO2 < 1100) {
                $("body").css("background-image", "linear-gradient(140deg, #212629, #395467)")
            } else {
                $("body").css("background-image", "linear-gradient(140deg, #765200, #FF0000)")
            }
        });
    }

    $(document).ready(function () {

        let item = $('#co2');
        let paramsCo2 = {
            initialValue: 0,
            higherValue: 3000,
            title: 'carbon dioxide in PPM',
            subtitle: '0'
        };

        let gauge = new GaugeChart(item, paramsCo2);
        gauge.init();
        //---
        // $('.gauge').each(function(index, item){
        //     let params = {
        //         initialValue: 780,
        //         higherValue: 1560,
        //         title: `Temperature ${index + 1}`,
        //         subtitle: '780 ºC'
        //     };
        //
        //     let gauge = new GaugeChart(item, params);
        //     gauge.init();
        // });
        setTimeout(updateCo2Handler, 20);
        // $('#random').click(function() {
        //
        //     $('.gauge').each(function(index, item) {
        //         let gauge = $(item).dxCircularGauge('instance');
        //         let randomNum = Math.round(Math.random() * 1560);
        //         let gaugeElement = $(gauge._$element[0]);
        //
        //         gaugeElement.find('.dxg-title text').last().html(`${randomNum} ºC`);
        //         gauge.value(randomNum);
        //     });
        // });
    });

});