//index.js
//获取应用实例
const app = getApp()

Page({
  data: {
    storyList: []
  },
  //事件处理函数
  goStory: function(event) {
      var id = event.currentTarget.id;
      wx.navigateTo({
        url: '../story/play?id='+id,
      });
  },
  onLoad: function () {
    var that = this;
      this.listStory(function(storyList){
        that.setData(
          {
            "storyList":storyList
          }
        )
      });
  },
  listStory:function(callback){
    wx.request({
      url: app.host+'/list',
      success:function(resp){
        var respData = resp.data;
        if(respData.code=="200"){
          callback(respData.data)
        }
      }
    })
  }
})
