const Vuego = {};

// Event bus implementation
Vuego.Bus = {
    _Read: null,
    _ReadQueue: [],
    _Write: null,

    connect: function() {
        const instance = this;

        // Connect Bus.Read endpoint
        let wsRead = new WebSocket('ws://' + window.location.host + '/bus', 'Bus.Read');
        wsRead.onmessage = function (ev) { instance._ReadQueue.push(JSON.parse(ev.data)) };
        wsRead.onopen = function (ev) { instance._Read = wsRead };
        wsRead.onclose = function (ev) { instance._Read = null };
        wsRead.onerror = function (ev) {};

        // Connect Bus.Write endpoint.
        let wsWrite = new WebSocket('ws://' + window.location.host + '/bus', 'Bus.Write');
        wsWrite.onmessage = function (ev) {};
        wsWrite.onopen = function (ev) { instance._Write = wsWrite };
        wsWrite.onclose = function (ev) { instance._Write = null };
        wsWrite.onerror = function (ev) {};
    },

    send: function(data) {
        const instance = this;

        // Send when it is possible
        var timerId = setTimeout(function tick() {
            if (instance._Write) {
                instance._Write.send(JSON.stringify(data))
            } else {
                timerId = setTimeout(tick, 1);
            }
        }, 0);
    },

    receive: function() {
        return this._ReadQueue.shift();
    }
};

// Event processing facilities
Vuego.Event = {
    serialize: function(e) {
        return e;
    }
};

// Action performing facilities
Vuego.Action = {
    perform: function(command, args) {
        if (Vuego.Action.hasOwnProperty(command)) {
            Vuego.Action[command](args);
        }
    },

    Say: function(message) {
        alert(message);
    }
};
