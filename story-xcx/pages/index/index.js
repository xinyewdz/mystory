//index.js
//获取应用实例
const app = getApp()

Page({
  data: {
    storyList: []
  },
  //事件处理函数
  goStory: function(event) {
      var id = event.currentTarget.id
      wx.navigateTo({
        url: '../story/play?id='+id,
      })
  },
  onLoad: function () {
      var storyList = [
        {
          "name":"test",
          "id":"1"
        }
      ]
      this.setData(
        {
          "storyList":storyList
        }
      )
  }
})
