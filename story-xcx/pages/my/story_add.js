// pages/my/my.js
const app = getApp()
const audio = wx.createInnerAudioContext();
Page({
  data: {
    story:{
      "name":"",
      "imageUrl":"",
      "audioUrl":"",
      "isPublic":false
    },
    showImgUpload:true
    
  },
  onLoad: function (options) {
    let id = options.id
    if(id&&id!=""){
      this.setData({
        showImgUpload:false
      })
      this.getStory(id)
    }
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
    app.postData("/save",that.data.story,function(){
      wx.switchTab({
        url: '../index/index',
        success:function(){
          var page = getCurrentPages().pop()
          if(page==undefined||page==null) return;
          page.onLoad();
        }
      })
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
  getStory:function(id){
    let req={
      "id":id
    }
    let that = this;
    app.postData("/story",req,function(resp){
      console.log(resp)
      that.setData({
        story:resp
      })
    });
  },
  nameEvent:function(event){
    var val = event.detail.value;
    console.log("name event."+val);
    let story = this.data.story
    story["name"]=val
    this.setData({
      story:story
    })
  },
  bindSwitchEvent:function(e){
    let val=e.detail.value;
    let story = this.data.story
    story["isPublic"]=val
    this.setData({
      story:story
    })
  },
  chooseAudio:function(){
    var that = this;
    wx.chooseMessageFile({
        count:1,
        type:"all",
        success (res) {
          const tempFilePaths = res.tempFiles;
          that.upload(tempFilePaths[0]["path"],that.data.story["name"],function(apath){
            let story = that.data.story;
            story["audioUrl"]=apath
            that.setData({
              story:story
            })
          });
        }
      })
  },
  chooseImg:function(){
    var that = this;
    let story = that.data.story;
    that.setData({
      showImgUpload:false
    });
    wx.chooseImage({
      success (res) {
        const tempFilePaths = res.tempFilePaths;
        that.upload(tempFilePaths[0],that.data.story["name"],function(apath){
          story["imageUrl"]=apath
          that.setData({
            showImgUpload:false,
            story:story
          })
        });
      }
    })
  }

})