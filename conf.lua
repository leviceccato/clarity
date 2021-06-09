local config = {
    identity = nil, -- The name of the save directory (string)
    appendidentity = false, -- Search files in source directory before save directory (boolean)
    version = '11.3', -- The LÖVE version this game was made for (string)
    console = false, -- Attach a console (boolean, Windows only)
    accelerometerjoystick = true, -- Enable the accelerometer on iOS and Android by exposing it as a Joystick (boolean)
    externalstorage = false, -- True to save files (and read from the save directory) in external storage on Android (boolean) 
    gammacorrect = false, -- Enable gamma-correct rendering, when supported by the system (boolean)
    audio = {
        mic = false, -- Request and use microphone capabilities in Android (boolean)
        mixwithsystem = true -- Keep background music playing when opening LOVE (boolean, iOS and Android only)
    },
    window = {
        title = 'Clarity', -- The window title (string)
        icon = 'images/tree-icon.png', -- Filepath to an image to use as the window's icon (string)
        width = 800, -- The window width (number)
        height = 600, -- The window height (number)
        borderless = false, -- Remove all border visuals from the window (boolean)
        resizable = true, -- Let the window be user-resizable (boolean)
        minwidth = 1, -- Minimum window width if the window is resizable (number)
        minheight = 1, -- Minimum window height if the window is resizable (number)
        fullscreen = false, -- Enable fullscreen (boolean)
        fullscreentype = 'desktop', -- Choose between "desktop" or "exclusive" fullscreen mode (string)
        vsync = 1, -- Vertical sync mode (number)
        msaa = 0, -- The number of samples to use with multi-sampled antialiasing (number)
        depth = nil, -- The number of bits per sample in the depth buffer
        stencil = nil, -- The number of bits per sample in the stencil buffer
        display = 1, -- Index of the monitor to show the window in (number)
        highdpi = false, -- Enable high-dpi mode for the window on a Retina display (boolean)
        usedpiscale = true, -- Enable automatic DPI scaling when highdpi is true as well (boolean)
        x = nil, -- The x-coordinate of the window's position in the specified display (number)
        y = nil, -- The y-coordinate of the window's position in the specified display (number)
    },
    modules = {
        audio = true, -- Enable the audio module (boolean)
        data = true, -- Enable the data module (boolean)
        event = true, -- Enable the event module (boolean)
        font = true, -- Enable the font module (boolean)
        graphics = true, -- Enable the graphics module (boolean)
        image = true, -- Enable the image module (boolean)
        joystick = true, -- Enable the joystick module (boolean)
        keyboard = true, -- Enable the keyboard module (boolean)
        math = true, -- Enable the math module (boolean)
        mouse = true, -- Enable the mouse module (boolean)
        physics = true, -- Enable the physics module (boolean)
        sound = true, -- Enable the sound module (boolean)
        system = true, -- Enable the system module (boolean)
        thread = true, -- Enable the thread module (boolean)
        timer = true, -- Enable the timer module (boolean), Disabling it will result 0 dt in update.
        touch = true, -- Enable the touch module (boolean)
        video = true, -- Enable the video module (boolean)
        window = true -- Enable the window module (boolean)
    }
}

love.conf = function()
    return config
end

return config