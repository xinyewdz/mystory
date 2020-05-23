//index.js
//获取应用实例
const app = getApp()
const adm = wx.getBackgroundAudioManager();
Page({
  data: {
    storyList: [],
    idx:0
  },
  
  onShow: function () {
    var flag = app.auth()
    if (!flag){
      wx.switchTab({
        url: '../my/my',
      })
    }
    var that = this;
    this.listStory(function(storyList){
      that.setData(
        {
          "storyList":storyList
        }
      );
      wx.setStorage({
        data: storyList,
        key: 'storyList',
      });
    });
  },
  next:function(){
    var idx = this.data.idx;
    var story = this.data.storyList[idx];
    idx++;
    if(idx>story.length){
      idx = 0;
    }
    
  },
  play:function(event){
    var id = event.currentTarget.id;
    var idx = event.currentTarget.dataset.idx;
    var story = this.data.storyList[idx];
    this.setData({
      idx:idx
    })
    adm["src"] = story.audioUrl
    adm["title"] = story.name
    adm["coverImgUrl"]= story.imageUrl
    adm.play()
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
