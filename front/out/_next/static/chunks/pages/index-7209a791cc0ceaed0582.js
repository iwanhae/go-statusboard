(self.webpackChunk_N_E=self.webpackChunk_N_E||[]).push([[405],{2562:function(e,n,t){"use strict";t.r(n);var r=t(5893),o=t(809),c=t.n(o),i=t(266),a=t(4468),s=t(9008),d=t(4476),l=t.n(d),u=t(7294),_=t(1870),f=t(9669),m=t.n(f);n.default=function(){var e=(0,u.useState)({}),n=e[0],t=e[1];(0,u.useEffect)((function(){var e=new EventSource("stream");return e.onopen=function(){console.log("Stream Open")},e.onmessage=function(e){var r=JSON.parse(e.data);r.checked_at=new Date(r.checked_at),void 0==n[r.name]&&(n[r.name]=[]),n[r.name]=[].concat((0,a.Z)(n[r.name]),[r]),t(Object.assign({},n))},e.onerror=function(e){console.error("Stream Error",e)},function(){e.close()}}),[]);var o=(0,u.useState)({}),d=o[0],f=o[1];(0,u.useEffect)((function(){(function(){var e=(0,i.Z)(c().mark((function e(){var n,t,r;return c().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.prev=0,e.next=3,m().get("/meta");case 3:n=e.sent,t=n.data,r={},t.forEach((function(e){r[e.name]=e})),f(r),console.log(r),e.next=14;break;case 11:e.prev=11,e.t0=e.catch(0),console.log(e.t0);case 14:case"end":return e.stop()}}),e,null,[[0,11]])})));return function(){return e.apply(this,arguments)}})()()}),[]);var h=Object.keys(d).map((function(e){var t,o=d[e],c=null===(t=n[e])||void 0===t?void 0:t.sort((function(e,n){return e.checked_at<n.checked_at?-1:1})),i=null===c||void 0===c?void 0:c.filter((function(e){return e.is_success})),a=null===c||void 0===c?void 0:c.filter((function(e){return!e.is_success})),s=null===i||void 0===i?void 0:i.map((function(e,n){return{x:e.checked_at.getTime(),y:e.duration}})),u=null===a||void 0===a?void 0:a.map((function(e,n){return{x:e.checked_at.getTime(),y:e.duration}}));return(0,r.jsxs)("div",{className:l().card,children:[(0,r.jsx)("h2",{children:o.name}),(0,r.jsx)("p",{children:o.description}),(null===c||void 0===c?void 0:c.length)>0&&(0,r.jsxs)("p",{className:l().fadeout,children:[c[0].checked_at.toLocaleTimeString()," ~ ",c[c.length-1].checked_at.toLocaleTimeString()]}),(0,r.jsx)("div",{children:(0,r.jsxs)(_.dp,{height:250,width:682,children:[(0,r.jsx)(_.pW,{}),(0,r.jsx)(_.xL,{}),(0,r.jsx)(_.B2,{title:"ms"}),(0,r.jsx)(_.eh,{data:s,color:"green"}),(0,r.jsx)(_.eh,{data:u,color:"red"})]})}),(null===a||void 0===a?void 0:a.length)>0&&(0,r.jsxs)("p",{children:["Last Failiure:",JSON.stringify(a[a.length-1].result)]})]},o.name)}));return(0,r.jsxs)("div",{className:l().container,children:[(0,r.jsxs)(s.default,{children:[(0,r.jsx)("title",{children:"Create Next App"}),(0,r.jsx)("meta",{name:"description",content:"Generated by create next app"}),(0,r.jsx)("link",{rel:"icon",href:"/favicon.ico"})]}),(0,r.jsxs)("main",{className:l().main,children:[(0,r.jsx)("h1",{className:l().title,children:"STATUSBOARD"}),(0,r.jsx)("div",{className:l().grid,children:h})]})]})}},5301:function(e,n,t){(window.__NEXT_P=window.__NEXT_P||[]).push(["/",function(){return t(2562)}])},4476:function(e){e.exports={fadeout:"Home_fadeout__2jQec",container:"Home_container__1EcsU",main:"Home_main__1x8gC",footer:"Home_footer__1WdhD",title:"Home_title__3DjR7",description:"Home_description__17Z4F",code:"Home_code__axx2Y",grid:"Home_grid__2Ei2F",card:"Home_card__2SdtB",logo:"Home_logo__1YbrH"}}},function(e){e.O(0,[774,119,888,179],(function(){return n=5301,e(e.s=n);var n}));var n=e.O();_N_E=n}]);