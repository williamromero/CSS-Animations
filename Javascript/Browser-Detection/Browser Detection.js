<script type="text/javascript">
    	// Opera 8.0+
	var isOpera = (!!window.opr && !!opr.addons) || !!window.opera || navigator.userAgent.indexOf(' OPR/') >= 0;
	// Firefox 1.0+
	var isFirefox = typeof InstallTrigger !== 'undefined';
	// At least Safari 3+: "[object HTMLElementConstructor]"
	var isSafari = Object.prototype.toString.call(window.HTMLElement).indexOf('Constructor') > 0;
	// Internet Explorer 6-11
	var isIE = /*@cc_on!@*/false || !!document.documentMode;
	// Edge 20+
	var isEdge = !isIE && !!window.StyleMedia;
	// Chrome 1+
	var isChrome = !!window.chrome && !!window.chrome.webstore;
	// Blink engine detection
	var isBlink = (isChrome || isOpera) && !!window.CSS;

		if ( isIE == true ){
			document.querySelector('.slider').style.display = 'none';
		}
		// var output = 'Detecting browsers by ducktyping:<hr>';
		// output += 'isFirefox: ' + isFirefox + '<br>';
		// output += 'isChrome: ' + isChrome + '<br>';
		// output += 'isSafari: ' + isSafari + '<br>';
		// output += 'isOpera: ' + isOpera + '<br>';
		// output += 'isIE: ' + isIE + '<br>';
		// output += 'isEdge: ' + isEdge + '<br>';
		// output += 'isBlink: ' + isBlink + '<br>';
				// document.body.innerHTML = output;
</script>
