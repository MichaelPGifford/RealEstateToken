var faded = false
$(document).ready(function(){
    let signInContainer = $(".sign-in-container")
    let closeBtn = $(".closeBtn")
    let overlay = $(".overlay")
    let signUpPopupBtn = $(".signUpPopupBtn")
    let signInPopupBtn = $(".signInPopupBtn")
    let signUpBtn = $(".signUpBtn")
    let signInBtn = $(".signInBtn")
    let sendEthBtn = $(".sendEthBtn")

    loginCheck()
    
    $('.dropdown-trigger').dropdown({
        "hover":true,
        
    })
    
    let clipboard = new ClipboardJS('.ethAddress')
    clipboard.on('success', function(e) {
        setTooltip('Copied!');
        hideTooltip();
      });
      
      clipboard.on('error', function(e) {
        setTooltip('Failed!');
        hideTooltip();
      });
    
    $('.tabs').tabs();

    $(signInContainer).click(signInBtnClick)
    $(closeBtn).click(fadeOut)
    $(overlay).click(fadeOut)
    $(signUpPopupBtn).click(signUpPopup)
    $(signInPopupBtn).click(signInPopup)
    $(signUpBtn).click(signUp)
    $(signInBtn).click(signIn)
    $(sendEthBtn).click(sendEth)
})

function signInBtnClick(){
    let popup = $(".overlay")
    let signInPopup = $(".signInPopup")
    popup.fadeIn(500)
    signInPopup.fadeIn(500)
}
function fadeOut(){
    let popup = $(".overlay")
    let signInPopup = $(".signInPopup")
    
    popup.fadeOut(700)
    signInPopup.fadeOut(200)
    
    
}
function signUpPopup(){
    $.when(setTimeout(function(){
        $(".signUpForm").fadeIn(400)
    }, 300)
    ).then($(".signInForm").fadeOut(300))
    
}
function signInPopup(){
    $.when(setTimeout(function(){
        $(".signInForm").fadeIn(400)
    }, 300)
    ).then($(".signUpForm").fadeOut(300))
    
}

function signUp(){
    data = {
        FName:$("#firstName").val(),
        LName:$("#lastName").val(),
        Email:$("#email1").val(),
        Password:$("#password1").val(),
        CPassword:$("#c_password1").val(),
    }
    $.ajax({
        type:"POST",
        url: "/api/sign-up",
        contentType: "application/json",
        data: JSON.stringify(data)
    })
}
function signIn(){
    data = {
        FName:"",
        LName:"",
        Email:$("#username").val(),
        Password:$("#password").val(),
        CPassword:"",
    }
    $.ajax({
        type:"POST",
        url: "/api/sign-in",
        contentType: "application/json",
        data: JSON.stringify(data)
    })
}

function loginCheck(){
    
    $.ajax({
        type:"GET",
        url: "/api/verify-id",
        contentType: "application/json",
        success: function(res){
            if(res){
                $(".sign-in-container").hide()
                $(".profile-container").show()
                res = JSON.parse(res)
                $($(".profile").find("i"))[0].after("hi, "+res.info.fname)
                $('.firstDropdownElm').after("HI, "+res.info.fname.toUpperCase())

                showWalletBalance(res)
                console.log(res.Wallet.PublicAddress)
            }else{
                console.log("gay")
            }
            
        }
    })
}

function showProfile(){

}
function showWalletBalance(res){
    
    $(".ethAddress").html(res.Wallet.PublicAddress + '<button data-toggle="tooltip" data-placement="bottom" ><i style="vertical-align:center;" class="tiny material-icons"> assignment</i></button>')
    $(".ethAddress").attr("data-clipboard-text", res.Wallet.PublicAddress)
    $.ajax({
        type:"GET",
        url: "/api/ether-balance",
        contentType: "application/json",
        success: function(res1){
            res1 = res1.replace(/"/g, "");
            console.log(res1)
            $('.ethbalance').html(res1+ " ETH")
            
        }
    })
}

function setTooltip(message) {
    $('.ethAddress').tooltip('hide')
      .attr('data-original-title', message)
      .tooltip('show');
  }
  
  function hideTooltip() {
    setTimeout(function() {
      $('.ethAddress').tooltip('hide');
    }, 1000);
  }

function sendEth(){
    data = {
        to: $("#toEthAddress").val(),
        amount: $("#ethAmountToSend").val()
    }
    $.ajax({
        type:"POST",
        url: "/api/send-ether",
        data: JSON.stringify(data),
        contentType: "application/json",
        success: function(res){
           
            console.log(res)
            $(".sendEthBtn").addClass("disabled")
            location.reload()
        }
    })
}