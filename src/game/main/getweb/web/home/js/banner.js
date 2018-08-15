/**
 *
 * 文件功能介绍
 *
 * @FileName: index.js
 * @Auther: Pandao
 * @QQ: 272383090
 * @CreateTime: 2013-02-20 10:00:20 
 * @UpdateTime: 2013-02-20 10:00:27 
 * Copyright@2013 泉州市科伟信息咨询有限公司版权所有，禁止非法用于商业用途，否则后果自负!
 */
/*
var bannerIndex = 0;
var bannerTotal, bannerTimer, index_banner_list;

function bannerListSlide() 
{
    index_banner_list.stop();
    index_banner_list.animate({marginLeft : '-'+(bannerIndex * 960)+'px'}, 1000, 'easeOutQuart');
	$("#banner_nav a").eq(bannerIndex).addClass("this").siblings().removeClass("this");
}

function nextBanner()
{
    if(bannerIndex < bannerTotal - 1) bannerIndex += 1;
    else bannerIndex = 0;
	bannerListSlide();
}

function prevBanner()
{ 
    if(bannerIndex > 0) bannerIndex -= 1;
    else bannerIndex = bannerTotal - 1;
	bannerListSlide();
}

function bannerAutoScroll()
{
    bannerTimer = setInterval(nextBanner, 6000);
}*/


var banner_index = 0, old_banner_index = banner_index, bannerTotal, bannerTimer, index_banner_list;

function banner_show() 
{ 
    $("#index_banner").stop();  
	index_banner_list.eq(old_banner_index).fadeOut(500);
	index_banner_list.eq(banner_index).delay(300).fadeIn(500);
	$("#banner_nav a").eq(banner_index).addClass("this").siblings().removeClass("this");
}

function next_banner()
{
	old_banner_index = banner_index;
    if(banner_index < bannerTotal - 1) banner_index += 1;
    else banner_index = 0;
	banner_show();
}

function prev_banner()
{ 
	old_banner_index = banner_index;
    if(banner_index > 0) banner_index -= 1;
    else banner_index = bannerTotal - 1;
	banner_show();
}

function banner_auto_show()
{
    bannerTimer = setInterval(next_banner, 5000);
}

//======================= Pro Start ==============================  

var pro_index = 0, pro_timer, pro_list, pro_total, pro_show_num = 4, pro_list_nav;

function pro_slider() { 
	pro_list_nav.find("a").eq(pro_index).addClass("this").siblings().removeClass("this");
	pro_list.stop().animate({marginLeft : "-"+(pro_index * 242)+"px"}, 1000, "easeOutQuart");
}

function pro_prev_action() {  

	if (pro_index > 0) {
		pro_index -= 1;
		pro_slider();
	}
	else{
		//pro_index = pro_total - 1;
	}
}

function pro_next_action() { 

	if (pro_index < pro_total - 4) {
		pro_index += 1; 
		pro_slider();
	}
}

function pro_auto_next() { 
	if (pro_index < pro_total - 4) pro_index += 1;  
	else pro_index = 0;
	pro_slider();
}

function pro_auto_slide() {
	pro_timer = setInterval(pro_auto_next, 5000);
}

//======================= Pro End ==============================  

//======================= Cases Start ==============================  

var cases_list; 

//======================= Cases End ==============================  

$(function() {
/*
    //Banner
    index_banner_list = $('#index_banner_list');
    bannerTotal = $('#index_banner_list li').length;
    index_banner_list.width((bannerTotal * 960)+10); 

	var nav_html = '';

	for (var i=0; i<bannerTotal; i++) {
		var className = (i == 0) ? ' class="this"' : '';
		nav_html += '<a href="javascript:;"'+className+'>&nbsp;&nbsp;</a>';
	}

	$('#banner_nav').append(nav_html);

    bannerAutoScroll();

    $('#index_banner').hover(function() {
        $('#prev_btn, #next_btn').fadeIn();
        clearInterval(bannerTimer);
    }, function() {
        $('#prev_btn, #next_btn').fadeOut();
        bannerAutoScroll();
    });

    $('#prev_btn').click(prevBanner);
    $('#next_btn').click(nextBanner); 

	$("#banner_nav a").click(function() {
		bannerIndex = $(this).index();
		bannerListSlide();
	}); 
*/
	
    //Banner
    index_banner_list = $('#index_banner_list li');
    bannerTotal = index_banner_list.length; 

	$('#prev_btn').css("opacity", "0.5");
	$('#next_btn').css("opacity", "0.5");
		
	if(bannerTotal > 1) {

		var nav_html = '';

		for (var i=0; i<bannerTotal; i++) 
		{
			if(i != 0) index_banner_list.eq(i).hide();
			var className = (i == 0) ? ' class="this"' : '';
			nav_html += '<a href="javascript:;"'+className+'>&nbsp;&nbsp;</a>';
		}

		$('#banner_nav').append(nav_html);
		banner_auto_show();

		$('#index_banner').hover(function() {
			$('#prev_btn, #next_btn').fadeIn();
			clearInterval(bannerTimer);
		}, function() {
			$('#prev_btn, #next_btn').fadeOut();
			banner_auto_show();
		});

		$('#prev_btn').click(prev_banner);
		$('#next_btn').click(next_banner); 

		$("#banner_nav a").click(function() {
			old_banner_index = banner_index;
			banner_index = $(this).index(); 
			banner_show();
		}); 
	}

	pro_list = $("#index_pro_list");
	pro_total = $("#index_pro_list li").length;  
	pro_list.css({width : (pro_total * 242)+20+"px"});

	function pro_list_slide() {		
		pro_list.stop().animate({marginLeft : "-"+242+"px"}, 1000, "easeOutQuart", function() {
			$(this).css({marginLeft : 0}).find("li:first").appendTo(this);
		});
	}

	function pro_list_auto_slide() {		
		 pro_timer = setInterval(pro_list_slide, 5000);
	}
	
	if(pro_total > 4) {
		pro_list_auto_slide();

		$("#index_pro").append().hover(function() { 
			clearInterval(pro_timer);
		}, function() { 
			pro_list_auto_slide();
		});  

		/*pro_auto_slide();

		var pro_nav_html = '';

		for (var i = 0, len = Math.ceil(pro_total / pro_show_num); i < len; i++) {
			var className = (i == 0) ? ' class="this"' : '';
			pro_nav_html += '<a href="javascript:;"'+className+'>'+i+'</a>';
		}
		//'<div id="pro_list_nav">'+pro_nav_html+'</div>'
		$("#index_pro").append().hover(function() { 
			clearInterval(pro_timer);
		}, function() { 
			pro_auto_slide();
		});  

		pro_list_nav = $("#pro_list_nav");
		pro_list_nav.find("a").bind("click", function() {
			pro_index = $(this).index();
			pro_slider();
		});*/


	} 

	var cases_total = $("#cases_list li").length;

	if(cases_total > 1) {

		$("#cases_slider").PDSlider({
			index : 0, 
			auto :true,
			width : 230,
			height : 132,
			slide : "left",
			child : 'li',          //移动的对象组成员
			eachNum : 1, 
			easing : 'easeOutQuart' 
		});

		
	} else {
		
		$(".PDSliderPrev").hide();
		$(".PDSliderNext").hide();
	}

	$("#cases_list li a img").hover(function() {
		$(this).animate({marginTop : "-30px"}, 600, 'easeOutQuart');
	}, function() {
		$(this).animate({marginTop : 0}, 600, 'easeOutQuart');
	});
});