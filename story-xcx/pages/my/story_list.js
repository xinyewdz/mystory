// pages/my/story_list.js
const app = getApp()

Page({
  data: {
    storyList:[]
  },

  onLoad: function (options) {
    var that = this;
    this.storyList(function(storyList){
      that.setData(
        {
          "storyList":storyList
        }
      )
    });
  },
  storyDelete:function(event){
      var id = event.currentTarget.id;
      console.log("remove "+id);
      var that = this;
      wx.request({
        url: app.host+'/remove?id='+id,
        success:function(resp){
          that.storyList(function(storyList){
            that.setData(
              {
                "storyList":storyList
              }
            )
          });
        }
      })
    
  },
  storyList:function(callback){
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