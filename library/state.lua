return function(arg)
    local s = {}
    s.worlds = {}
    s.activeWorld = nil
    s.controls = {}

    local keyMap = {
        ['space'] = 'jump',
        ['up'] = 'up',
        ['w'] = 'up',
        ['left'] = 'left',
        ['a'] = 'left',
        ['right'] = 'right',
        ['d'] = 'right',
        ['down'] = 'down',
        ['s'] = 'down',
        ['escape'] = 'menu',
        ['`'] = 'debug'
    }
    for _, control in pairs(keyMap) do
        s.controls[control] = false
    end

    s.updateControls = function(key, isDown)
        local control = keyMap[key]
        if control then
            s.controls[control] = isDown
        end
    end

    s.addWorld = function(world)
        s.worlds[world.name] = world
    end

    s.activateWorld = function(worldName)
        if s.activeWorld then
            s.activeWorld.exit()
        end
        s.worlds[worldName].load(arg)
        s.activeWorld = s.worlds[worldName]
    end

    s.update = function(dt)
        s.activeWorld.update(dt)
    end

    s.draw = function()
        s.activeWorld.draw()
    end

    s.keypressed = function(key)
        s.updateControls(key, true)
        s.activeWorld.keypressed(key)
    end

    s.keyreleased = function(key)
        s.updateControls(key, false)
        s.activeWorld.keyreleased(key)
    end

    s.resize = function(w, h)
        s.activeWorld.resize(w, h)
    end

    return s
end