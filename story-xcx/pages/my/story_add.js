// pages/my/my.js
const app = getApp()
const audio = wx.createInnerAudioContext();
Page({
  data: {
    "name":"",
    "image":"",
    "audioPath":"",
    "audioName":""
  },
  onLoad: function (options) {
    var token = app.getToken();
    if(!token){
      wx.switchTab({
        url: '../my/my',
      })
      return;
    }
  },
  saveStory:function(){
    wx.showLoading({
      "title":"上传中"
    });
    var that = this;
    let name = this.data.name;
    console.log("savestory,"+name);
    that.upload(that.data.audioPath,name,function(aPath){
      that.upload(that.data.image,name,function(iPath){
        var data = {
          name:name,
          audio:aPath,
          image:iPath
        };
        app.postData("/save",data,function(){
          wx.switchTab({
            url: '../index/index',
            success:function(){
              var page = getCurrentPages().pop()
              if(page==undefined||page==null) return;
              page.onLoad();
            }
          })
        });
      });
    });
    
  },
  upload:function(path,name,callback){
    var token = app.getToken();
    wx.uploadFile({
      filePath: path,
      name: "file",
      url: app.host+"/upload",
      header:{
        "token":token
      },
      formData:{
        "name":name
      },
      success:function(res,code){
        let resData = JSON.parse(res.data);
        console.log(resData)
        if(resData.code=="200"){
          callback(resData.data);
        }else{
          wx.showToast({
            title: resData.msg
          })
        }
        console.log("upload finish.code="+resData.code);
      },
      fail:function(){
        console.log("upload fail.path="+path);
      }
    })
  },
  nameEvent:function(event){
    var name = event.detail.value;
    console.log("name event."+name);
    this.setData({
      "name":name
    })
  },
  chooseAudio:function(){
    var that = this;
    console.log("chooseAudio event")
    wx.chooseMessageFile({
        count:1,
        type:"all",
        success (res) {
          const tempFilePaths = res.tempFiles;
          that.setData({
            "audioPath":tempFilePaths[0]["path"],
            "audioName":tempFilePaths[0]["name"]
          })
        }
      })
  },
  chooseImg:function(){
    var that = this;
    console.log("chooseImg event")
    wx.chooseImage({
      success (res) {
        const tempFilePaths = res.tempFilePaths;
        that.setData({
          "image":tempFilePaths[0]
        })
      }
    })
  }

})