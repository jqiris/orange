cc.Class({
    extends: cc.Component,

    properties: {
        // foo: {
        //    default: null,      // The default value will be used only when the component attaching
        //                           to a node for the first time
        //    url: cc.Texture2D,  // optional, default is typeof default
        //    serializable: true, // optional, default is true
        //    visible: true,      // optional, default is true
        //    displayName: 'Foo', // optional
        //    readonly: false,    // optional, default is false
        // },
        // ...
        _lblTip: null,
        _lastPing: 0,
    },

    // use this for initialization
    onLoad: function() {
        var self = this;
        var try_time = 0;
        var fnTestServerOn = function() {
            cc.vv.net.ping(function(ret) {
                if (ret) {
                    try_time = 0;
                    cc.vv.gameNetMgr.reset();
                    //cc.director.loadScene('hall');
                    var roomId = cc.vv.userMgr.oldRoomId;
                    if (roomId != null) {
                        cc.vv.userMgr.oldRoomId = null;
                        cc.vv.userMgr.enterRoom(roomId, function(ret) {
                            if (ret.errcode != 0) {
                                cc.vv.gameNetMgr.roomId = null;
                                cc.director.loadScene('hall');
                            }
                        });
                    }
                } else if (try_time < 5) {
                    try_time += 1
                    setTimeout(fnTestServerOn, 3000);
                } else {
                    try_time = 0
                    cc.vv.gameNetMgr.roomId = null;
                    cc.director.loadScene('hall');
                    cc.vv.wc.hide();
                }
            });
        }

        var fn = function(data) {
            self.node.off('disconnect', fn);
            cc.vv.wc.show("正在重连...");
            fnTestServerOn();
        };
        console.log("adasfdasdfsdf");

        this.node.on('login_finished', function() {
            cc.vv.wc.hide();
            self.node.on('disconnect', fn);
        });
        this.node.on('disconnect', fn);
    },
    // called every frame, uncomment this function to activate update callback
    update: function(dt) {},
});