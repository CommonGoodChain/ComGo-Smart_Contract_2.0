## ComGo-Smart_Contract_2.0

Please visit the [installation instructions](http://hyperledger-fabric.readthedocs.io/en/latest/install.html)
to ensure you have the correct prerequisites installed. Please use the
version of the documentation that matches the version of the software you
intend to use to ensure alignment.

## Download Binaries and Docker Images

The installation instructions will utilize `scripts/bootstrap.sh` (available in the fabric repository)
script to download all of the requisite Hyperledger Fabric binaries and docker
images, and tag the images with the 'latest' tag. Optionally,
specify a version for fabric, fabric-ca and thirdparty images. If versions
are not passed, the latest available versions will be downloaded.

The script will also clone fabric-samples repository using the version tag that
is aligned with the Fabric version.

You can also download the script and execute locally:

```bash
# Fetch bootstrap.sh from fabric repository using
curl -sS https://raw.githubusercontent.com/hyperledger/fabric/master/scripts/bootstrap.sh -o ./scripts/bootstrap.sh
# Change file mode to executable
chmod +x ./scripts/bootstrap.sh
# Download binaries and docker images
./scripts/bootstrap.sh [version] [ca version] [thirdparty_version]
```
