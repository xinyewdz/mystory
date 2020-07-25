//index.js
//获取应用实例
const app = getApp()
const adm = app.adm;
Page({
  data: {
    storyList: [],
    idx:0,
    playStory:{},
    showPlayer:false,
    playState:0
  },
  onShareAppMessage:function(){

  },
  onReady:function(){
    adm.onEnded(this.next);
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
    var storyList = this.data.storyList;
    idx++;
    if(idx>=storyList.length){
      idx = 0;
    }
    var story = storyList[idx];
    this.setData({
      playStory:story
    })
    this.playStory(story,idx)

  },
  play:function(event){
    var id = event.currentTarget.id;
    var idx = event.currentTarget.dataset.idx;
    var story = this.data.storyList[idx];
    this.setData({
      showPlayer:true,
      playState:1,
      playStory:story
    })
    this.playStory(story,idx)
  },
  playStory:function(story,idx){
    this.setData({
      idx :idx
    })
    adm["src"] = story.audioUrl
    adm["title"] = story.name
    adm["coverImgUrl"]= story.imageUrl
    adm.play()
    var req={
      "storyId":story.id
    }
    app.postData("/play",req,function(){

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
    app.postData("/play/list",{},function(respData){
      callback(respData)
    });
  },
  handPlayEvent:function(){
    let playState = this.data.playState;
    if(playState==1){
      this.setData({
        playState:2,
      });
      adm.pause();
    }else{
      this.setData({
        playState:1,
      });
      adm.play()
    }
  }
})
