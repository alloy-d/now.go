function popAndUpdate() {
    var doneRequest = new Request({
            'url': '/done',
            'method': 'POST',
            'data': 'done',
            });
    var nextRequest = new Request.JSON({
            'url': '/now.json',
            'method': 'GET',
            });
    nextRequest.addEvent('success', function (text) {
            $("theThing").set('text', text);
            });
    doneRequest.addEvent('complete', function () {
            nextRequest.send();
            });

    doneRequest.send();
}

function showDoneButton() {
    var doneButton = new Element('p', {
            'id': 'theDoneButton',
            'html': 'done.',
            'events': {
                'click': function () {
                    popAndUpdate();
                    hideDoneButton();
                },
                'mouseout': function () {
                    hideDoneButton();
                },
            },
            'styles': {
                'font-size': $('theThing').getStyle('height'),
                'height': $('theThing').getStyle('height'),
                'width': $('theThing').getStyle('width'),
            },
            'tween': {duration: 200},
        });
    doneButton.fade('hide');
    doneButton.setPosition($('theThing').getPosition());
    doneButton.inject($('box'), 'bottom');
    doneButton.fade(0.95);
}

function hideDoneButton() {
    $('theDoneButton').get('tween', {
        duration: 300,
        property: 'opacity',
    }).addEvent('complete', function(){$('theDoneButton').destroy();}).start(0);
}

window.addEvent('domready', function () {
        $("theThing").addEvent('mouseover', showDoneButton);
        });
