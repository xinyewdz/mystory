//index.js
//获取应用实例
const app = getApp()

Page({
  data: {
    storyList: []
  },
  
  onShow: function () {
    var token = app.getToken();
    if(!token){
      console.log("token empty")
      wx.switchTab({
        url: '../my/my',
      })
      return;
    }
    var that = this;
    this.listStory(function(storyList){
      that.setData(
        {
          "storyList":storyList
        }
      )
    });
  },
  //事件处理函数
  goStory: function(event) {
    var id = event.currentTarget.id;
    wx.navigateTo({
      url: '../story/play?id='+id,
    });
},
  listStory:function(callback){
    app.postData("/list",{},function(respData){
      callback(respData)
    });
  }
})
