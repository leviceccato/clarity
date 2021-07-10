return function(arg)
    local s = {}
    s.worlds = {}
    s.activeWorlds = {}

    s.inputs = {}
    s.controls = {
        ['mouse1'] = 'click',
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
    for _, input in pairs(s.controls) do
        s.inputs[input] = nil
    end

    s.loadWorld = function(world)
        world.load(arg)
        s.worlds[world.name] = world
    end

    s.activateWorlds = function(worldNames)
        s.activeWorlds = worldNames
    end

    s.update = function(dt)
        for index = 1, #s.activeWorlds do
            s.worlds[s.activeWorlds[index]].update(dt)
        end
    end

    s.draw = function()
        for index = 1, #s.activeWorlds do
            s.worlds[s.activeWorlds[index]].draw()
        end
    end

    s.exit = function()
        for index = 1, #s.activeWorlds do
            s.worlds[s.activeWorlds[index]].exit()
        end
    end

    s.mousepressed = function(x, y, button, isTouch, pressCount)
        for index = 1, #s.activeWorlds do
            s.worlds[s.activeWorlds[index]].mousepressed(x, y, button, isTouch, pressCount)
        end
    end

    s.mousereleased = function(button)
        for index = 1, #s.activeWorlds do
            s.worlds[s.activeWorlds[index]].mousereleased(button)
        end
    end

    s.keypressed = function(key, isRepeat)
        for index = 1, #s.activeWorlds do
            s.worlds[s.activeWorlds[index]].keypressed(key, isRepeat)
        end
    end

    s.keyreleased = function(key)
        for index = 1, #s.activeWorlds do
            s.worlds[s.activeWorlds[index]].keyreleased(key)
        end
    end

    s.resize = function(w, h)
        for index = 1, #s.activeWorlds do
            s.worlds[s.activeWorlds[index]].resize(w, h)
        end
    end

    return s
end