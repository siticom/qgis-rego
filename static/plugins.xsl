<?xml version="1.0" encoding="ISO-8859-1"?>
<!-- Inspired by https://gitlab.com/GIS-projects/phpQGISrepository -->
<xsl:stylesheet version="1.0"
    xmlns:xsl="http://www.w3.org/1999/XSL/Transform">

    <xsl:template match="plugins">
        <html>
            <head>
                <meta name="generator" content="Golang QGIS Repository" />
                <title>QGIS Plugins Repository</title>
                <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.1/css/bootstrap.min.css" integrity="sha384-WskhaSGFgHYWDcbwN70/dfYBj47jz9qbsMId/iRN3ewGhXQFZCSftd1LZCfmhktB" crossorigin="anonymous" />
                <style>
                .footer {
                    width: 100%;
                    height: 60px; 
                    line-height: 60px; 
                    background-color: #f5f5f5;
                }
                </style>
            </head>
            <body>
                <div class="jumbotron">
                    <h1>
                        <div class="d-none d-md-block float-left mr-5">
                            <img width="64" height="64" src="data:image/png;base64,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" />
                        </div>QGIS Plugins Repository</h1>
                </div>
                <div class="container-fluid">
                    <div class="row">
                        <div class="col-lg-3 col-md-4">
                            <div class="list-group">
                                <div class="text-white bg-dark list-group-item">
                                    <h2>Plugins</h2>
                                </div>
                                <xsl:for-each select="/plugins/pyqgis_plugin">
                                    <xsl:element name="a">
                                        <xsl:attribute name="href">
                                            <xsl:value-of select="concat('#',@id)" />
                                        </xsl:attribute>
                                        <xsl:attribute name="class">
                                list-group-item list-group-item-action list-group-item-light
                                        </xsl:attribute>
                                        <xsl:value-of select="@name" />
                                        <xsl:text> v</xsl:text>
                                        <xsl:value-of select="@version" />
                                    </xsl:element>
                                </xsl:for-each>
                            </div>
                        </div>
                        <div class="col">
                            <xsl:for-each select="/plugins/pyqgis_plugin">
                                <div class="card mt-0 mb-3">
                                    <div class="card-header">
                                        <xsl:element name="h2">
                                            <xsl:attribute name="id">
                                                <xsl:value-of select="@id" />
                                            </xsl:attribute>
                                            <xsl:element name="div">
                                                <xsl:attribute name="style">
                                                    <xsl:text>display: flex; align-items: center;</xsl:text>
                                                </xsl:attribute>
                                                <xsl:if test="icon_base64 != ''">
                                                    <xsl:element name="img">
                                                        <xsl:attribute name="src">
                                                            <xsl:text>data:image/png;base64,</xsl:text>
                                                            <xsl:value-of select="icon_base64" />
                                                        </xsl:attribute>
                                                        <xsl:attribute name="width">30</xsl:attribute>
                                                        <xsl:attribute name="height">30</xsl:attribute>
                                                        <xsl:attribute name="style">margin-right: 10px;</xsl:attribute>
                                                    </xsl:element>
                                                </xsl:if>
                                                <xsl:value-of select="@name" />
                                                <xsl:text> v</xsl:text>
                                                <xsl:value-of select="@version" />
                                            </xsl:element>
                                        </xsl:element>
                                    </div>
                                    <div class="card-body">
                                        <div class="card-title">
                                            <h3>
                                                <xsl:value-of select="description" />
                                            </h3>
                                        </div>
                                        <div class="about">
                                            <p>
                                                <i>
                                                    <xsl:value-of select="about" />
                                                </i>
                                            </p>
                                        </div>
                                        <div class="tags">
                                            <strong>Tags:</strong>&#160;<xsl:value-of select="tags" />
                                        </div>
                                        <div class="download">
                                            <strong>Download:</strong>&#160;
                                            <xsl:element name="a">
                                                <xsl:attribute name="href">
                                                    <xsl:value-of select="download_url" />
                                                </xsl:attribute>
                                                <xsl:value-of select="file_name" />
                                            </xsl:element>
                                        </div>
                                        <xsl:if test="author != ''">
                                            <strong>Author:</strong>&#160;
                                            <xsl:element name="a">
                                                <xsl:attribute name="href">
                                                    <xsl:text>mailto:</xsl:text>
                                                    <xsl:value-of select="email" />
                                                </xsl:attribute>
                                                <xsl:value-of select="author" />
                                            </xsl:element>
                                        </xsl:if>
                                        <div class="@version">
                                            <strong>Version:</strong>&#160;<xsl:value-of select="@version" />
                                        </div>
                                        <div class="experimental">
                                            <strong>Experimental:</strong>&#160;<xsl:value-of select="experimental" />
                                        </div>
                                        <div class="deprecated">
                                            <strong>Deprecated:</strong>&#160;<xsl:value-of select="deprecated" />
                                        </div>
                                        <div class="minversion">
                                            <strong>Minimum QGIS Version:</strong>&#160;<xsl:value-of select="qgis_minimum_version" />
                                        </div>
                                        <div class="maxversion">
                                            <strong>Maximum QGIS Version:</strong>&#160;<xsl:value-of select="qgis_maximum_version" />
                                        </div>
                                        <xsl:if test="homepage != ''">
                                            <div class="homepage">
                                                <strong>Homepage:</strong>&#160;
                                                <xsl:element name="a">
                                                    <xsl:attribute name="href">
                                                        <xsl:value-of select="homepage" />
                                                    </xsl:attribute>
                                                    <xsl:value-of select="homepage" />
                                                </xsl:element>
                                            </div>
                                        </xsl:if>
                                        <xsl:if test="tracker != ''">
                                            <div class="tracker">
                                                <strong>Tracker:</strong>&#160;
                                                <xsl:element name="a">
                                                    <xsl:attribute name="href">
                                                        <xsl:value-of select="tracker" />
                                                    </xsl:attribute>
                                                    <xsl:value-of select="tracker" />
                                                </xsl:element>
                                            </div>
                                        </xsl:if>
                                        <xsl:if test="repository != ''">
                                            <div class="repository">
                                                <strong>Repository:</strong>&#160;
                                                <xsl:element name="a">
                                                    <xsl:attribute name="href">
                                                        <xsl:value-of select="repository" />
                                                    </xsl:attribute>
                                                    <xsl:value-of select="repository" />
                                                </xsl:element>
                                            </div>
                                        </xsl:if>
                                    </div>
                                </div>
                            </xsl:for-each>

                        </div>
                    </div>
                </div>
                <footer class="footer">
                    <div class="container text-center">
                        <span class="text-muted">This plugin repository is generated using Golang QGIS Repository v0.1 (Interface by <a href="https://gitlab.com/GIS-projects/phpQGISrepository" target="_blank">phpQGISrepository</a>).
                        </span>
                    </div>
                </footer>
                <script src="https://code.jquery.com/jquery-3.3.1.slim.min.js" integrity="sha384-q8i/X+965DzO0rT7abK41JStQIAqVgRVzpbzo5smXKp4YfRvH+8abtTE1Pi6jizo" crossorigin="anonymous"></script>
                <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.3/umd/popper.min.js" integrity="sha384-ZMP7rVo3mIykV+2+9J3UJ46jBk0WLaUAdn689aCwoqbBJiSnjAK/l8WvCWPIPm49" crossorigin="anonymous"></script>
                <script src="https://stackpath.bootstrapcdn.com/bootstrap/4.1.1/js/bootstrap.min.js" integrity="sha384-smHYKdLADwkXOn1EmN1qk/HfnUcbVRZyYmZ4qpPea6sjB/pTJ0euyQp0Mk8ck+5T" crossorigin="anonymous"></script>
            </body>
        </html>

    </xsl:template>

</xsl:stylesheet>
