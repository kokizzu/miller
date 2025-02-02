..
    PLEASE DO NOT EDIT DIRECTLY. EDIT THE .rst.in FILE PLEASE.

Installation
================================================================

Prebuilt executables via package managers
----------------------------------------------------------------

`Homebrew <https://brew.sh/>`_ installation support for OSX is available via

.. code-block:: none
   :emphasize-lines: 1,1

    brew update && brew install miller

...and also via `MacPorts <https://www.macports.org/>`_:

.. code-block:: none
   :emphasize-lines: 1,1

    sudo port selfupdate && sudo port install miller

You may already have the ``mlr`` executable available in your platform's package manager on NetBSD, Debian Linux, Ubuntu Xenial and upward, Arch Linux, or perhaps other distributions. For example, on various Linux distributions you might do one of the following:

.. code-block:: none
   :emphasize-lines: 1,1

    sudo apt-get install miller

.. code-block:: none
   :emphasize-lines: 1,1

    sudo apt install miller

.. code-block:: none
   :emphasize-lines: 1,1

    sudo yum install miller

On Windows, Miller is available via `Chocolatey <https://chocolatey.org/>`_:

.. code-block:: none
   :emphasize-lines: 1,1

    choco install miller

Prebuilt executables via GitHub per release
----------------------------------------------------------------

Please see https://github.com/johnkerl/miller/releases where there are builds for OSX Yosemite, Linux x86-64 (dynamically linked), and Windows (via Appveyor build artifacts).

Miller is autobuilt for **Linux** using **Travis** on every commit (https://travis-ci.org/johnkerl/miller/builds). This was set up by the generous assistance of `SikhNerd <https://github.com/SikhNerd>`_ on Github, tracked in https://github.com/johnkerl/miller/issues/15. Analogously, Miller is autobuilt for **Windows** using the **Appveyor** continuous-build system: https://ci.appveyor.com/project/johnkerl/miller.

Miller releases from `5.1.0 <https://github.com/johnkerl/miller/releases/tag/v5.1.0w>`_ onward will have a precompiled Windows binary, in addition to the MacOSX and Linux 64-bit precompiled binaries as on previous releases.  Specifically, at https://ci.appveyor.com/project/johnkerl/miller you can select *Latest Build* and then *Artifacts* to always get the current head build. Miller releases from 5.3.0 onward will simply point to a particular Appveyor artifact associated with the release.

Building from source
----------------------------------------------------------------

Please see :doc:`build`.
