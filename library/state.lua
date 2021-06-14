return function()
    local s = {}
    s.worlds = {}
    s.activateWorld = nil
    s.controls = {}

    local controls = {
        'space' = 'jump',
        'up' = 'up',
        'w' = 'up',
        'left' = 'left'
        'a' = 'left',
        'right' = 'right',
        'd' = 'right'
        'down' = 'down'
        's' = 'down',
        'escape' = 'menu',
        '`' = 'debug'
    }
    for _, control in pairs(controls) do
        if ~s.controls[control] then
            s.controls[control] = false
        end
    end

    s.updateControls = function(key, isActive) do
        local control = controls[key]
        s.controls[control] = isActive
    end

    s.addWorld = function(world)
        s.worlds[world.name] = world
    end

    s.activateWorld = function(worldName)
        s.activateWorld = s.worlds[worldName]
    end

    s.updateControls = function()
    end

    s.load = function(arg)
        s.activateWorld.load(arg)
    end

    s.update = function(dt)
        s.activateWorld.update(dt)
    end

    s.draw = function()
        s.activateWorld.draw()
    end

    s.keypressed = function(key)
        s.updateControls(key, true)
        s.activateWorld.keypressed(key)
    end

    s.keyreleased = function(key)
        s.updateControls(key, false)
        s.activateWorld.keyreleased(key)
    end

    s.resize = function(...)
        s.activateWorld.resize(...)
    end

    return s
end