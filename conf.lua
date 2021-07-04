love.conf = function(c)
    c.identity = nil                    -- The name of the save directory (string)
    c.appendidentity = false            -- Search files in source directory before save directory (boolean)
    c.version = '11.3'                  -- The LÖVE version this game was made for (string)
    c.console = false                   -- Attach a console (boolean, Windows only)
    c.accelerometerjoystick = true      -- Enable the accelerometer on iOS and Android by exposing it as a Joystick (boolean)
    c.externalstorage = false           -- True to save files (and read from the save directory) in external storage on Android (boolean)
    c.gammacorrect = false              -- Enable gamma-correct rendering, when supported by the system (boolean)

    c.audio.mic = false                 -- Request and use microphone capabilities in Android (boolean)
    c.audio.mixwithsystem = true        -- Keep background music playing when opening LOVE (boolean, iOS and Android only)

    c.window.title = 'Clarity'          -- The window title (string)
    c.window.icon = nil                 -- Filepath to an image to use as the window's icon (string)
    c.window.width = 960                -- The window width (number)
    c.window.height = 540               -- The window height (number)
    c.window.borderless = false         -- Remove all border visuals from the window (boolean)
    c.window.resizable = true           -- Let the window be user-resizable (boolean)
    c.window.minwidth = 480             -- Minimum window width if the window is resizable (number)
    c.window.minheight = 270            -- Minimum window height if the window is resizable (number)
    c.window.fullscreen = false         -- Enable fullscreen (boolean)
    c.window.fullscreentype = 'desktop' -- Choose between "desktop" fullscreen or "exclusive" fullscreen mode (string)
    c.window.vsync = 1                  -- Vertical sync mode (number)
    c.window.msaa = 0                   -- The number of samples to use with multi-sampled antialiasing (number)
    c.window.depth = nil                -- The number of bits per sample in the depth buffer
    c.window.stencil = nil              -- The number of bits per sample in the stencil buffer
    c.window.display = 1                -- Index of the monitor to show the window in (number)
    c.window.highdpi = false            -- Enable high-dpi mode for the window on a Retina display (boolean)
    c.window.usedpiscale = true         -- Enable automatic DPI scaling when highdpi is set to true as well (boolean)
    c.window.x = nil                    -- The x-coordinate of the window's position in the specified display (number)
    c.window.y = nil                    -- The y-coordinate of the window's position in the specified display (number)

    c.modules.audio = true              -- Enable the audio module (boolean)
    c.modules.data = true               -- Enable the data module (boolean)
    c.modules.event = true              -- Enable the event module (boolean)
    c.modules.font = true               -- Enable the font module (boolean)
    c.modules.graphics = true           -- Enable the graphics module (boolean)
    c.modules.image = true              -- Enable the image module (boolean)
    c.modules.joystick = true           -- Enable the joystick module (boolean)
    c.modules.keyboard = true           -- Enable the keyboard module (boolean)
    c.modules.math = true               -- Enable the math module (boolean)
    c.modules.mouse = true              -- Enable the mouse module (boolean)
    c.modules.physics = true            -- Enable the physics module (boolean)
    c.modules.sound = true              -- Enable the sound module (boolean)
    c.modules.system = true             -- Enable the system module (boolean)
    c.modules.thread = true             -- Enable the thread module (boolean)
    c.modules.timer = true              -- Enable the timer module (boolean), Disabling it will result 0 delta time in love.update
    c.modules.touch = true              -- Enable the touch module (boolean)
    c.modules.video = true              -- Enable the video module (boolean)
    c.modules.window = true             -- Enable the window module (boolean)

    -- Custom config
    c.renderWidth = 480
    c.renderHeight = 270
    c.colFg = {215 / 255, 196 / 255, 91 / 255, 1}

    love.c = c
end