 <!-- Footer -->
<footer>
    <div class="container">
        <div class="row">
            <div class="col-lg-8 col-lg-offset-2 col-md-10 col-md-offset-1">
                <ul class="list-inline text-center">
                    {{if .Site.Twitter}}
                    <li>
                        <a href="{{.Site.Twitter}}">
                            <span class="fa-stack fa-lg">
                                <i class="fa fa-circle fa-stack-2x"></i>
                                <i class="fa fa-twitter fa-stack-1x fa-inverse"></i>
                            </span>
                        </a>
                    </li>
                    {{end}}
                    {{if .Site.Facebook}}
                    <li>
                        <a href="{{.Site.Facebook}}">
                            <span class="fa-stack fa-lg">
                                <i class="fa fa-circle fa-stack-2x"></i>
                                <i class="fa fa-facebook fa-stack-1x fa-inverse"></i>
                            </span>
                        </a>
                    </li>
                    {{end}}
                    {{if .Site.Github}}
                    <li>
                        <a href="{{.Site.Github}}">
                            <span class="fa-stack fa-lg">
                                <i class="fa fa-circle fa-stack-2x"></i>
                                <i class="fa fa-github fa-stack-1x fa-inverse"></i>
                            </span>
                        </a>
                    </li>
                    {{end}}
                </ul>
                <p class="copyright text-muted">Powered by <a href="https://github.com/cwen-coder/ljgo">ljgo</a> &nbsp;  </a> &copy; 2016 &nbsp;<a href="{{.Sijte.URL}}">{{.Site.Title}}</p>
            </div>
        </div>
    </div>
</footer>
<!-- jQuery -->
<script src="/static/vendor/jquery/jquery.min.js"></script>

<!-- Bootstrap Core JavaScript -->
<script src="/static/vendor/bootstrap/js/bootstrap.min.js"></script>

<!-- Contact Form JavaScript -->
<script src="/static/js/jqBootstrapValidation.js"></script>
<script src="/static/js/contact_me.js"></script>

<!-- Theme JavaScript -->
<script src="/static/js/clean-blog.js"></script>
