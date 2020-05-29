// pages/my/story_list.js
const app = getApp()

Page({
  data: {
    storyList:[]
  },

  onLoad: function (options) {
    var token = app.getToken();
    if(!token){
      wx.switchTab({
        url: '../my/my',
      })
      return;
    }
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
      var data = {
        "id":id
      }
      app.postData("/remove",data,function(){
        that.storyList(function(storyList){
          that.setData(
            {
              "storyList":storyList
            }
          )
        });
      })
    
  },
  storyList:function(callback){
    app.postData("/list",{},function(respData){
      callback(respData);
    });
  }
})