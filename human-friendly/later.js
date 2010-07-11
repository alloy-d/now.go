document.addEvent('domready', function () {
        $('busyIndicator').fade('hide');
        var addRequest = new Request({
            'url': '/later/',
            'method': 'POST',
            'link': 'chain',
            });
        addRequest.addEvent('success', function () {
            $('theThing').set('value', null);
            $('busyIndicator').fade('out');
            });
        addRequest.addEvent('failure', function () {
            $("busyIndicator").fade('out');
            $('theThing').blur();
            $$("input[type=submit]").set('value', 'x');
            $$("input").set('disabled', 'disabled');
            alert("Something went wrong!");
            });
        addRequest.addEvent('request', function () {
            $('busyIndicator').fade('in');
            });
        $('theForm').addEvent('submit', function () {
            addRequest.send('thing=' + $('theThing').get('value'));
            return false;
            });
        });
