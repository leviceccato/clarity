return function(name)
    local w = {}
    w.name = name
    w.systems = {}
    w.entities = {}

    w.addEntity = function(entity, shouldUpdate)
        shouldUpdate = shouldUpdate or false
        w.entities[#w.entities + 1] = entity
        if shouldUpdate then
            w.updateSystems()
        end
    end

    w.addSystem = function(system)
        w.systems[#w.systems + 1] = system
    end

    w.sortEntities = function()
        for index = 1, #w.systems do
            w.systems[index].sortEntities()
        end
    end

    w.updateSystems = function()
        local system
        local entity
        local component
        local hasComponents = true
        for systemIndex = 1, #w.systems do
            system = w.systems[systemIndex]
            for entityIndex = 1, #w.entities do
                entity = w.entities[entityIndex]
                for componentIndex = 1, #system.components do
                    component = system.components[componentIndex]
                    if entity[component] == nil then
                        hasComponents = false
                    end
                end
                if hasComponents then
                    system.addEntity(entity)
                end
                hasComponents = true
            end
        end
    end

    w.load = function(arg)
        for index = 1, #w.systems do
            w.systems[index].load(arg)
        end
    end

    w.exit = function()
        for index = 1, #w.systems do
            w.systems[index].exit()
        end
    end

    w.update = function(dt)
        for index = 1, #w.systems do
            w.systems[index].update(dt)
        end
    end

    w.draw = function()
        for index = 1, #w.systems do
            w.systems[index].draw()
        end
    end

    w.mousepressed = function(x, y, button, isTouch, pressCount)
        for index = 1, #w.systems do
            w.systems[index].keypressed(x, y, button, isTouch, pressCount)
        end
    end

    w.mousereleased = function(button)
        for index = 1, #w.systems do
            w.systems[index].mousereleased(button)
        end
    end

    w.keypressed = function(key, isRepeat)
        for index = 1, #w.systems do
            w.systems[index].keypressed(key)
        end
    end

    w.keyreleased = function(key)
        for index = 1, #w.systems do
            w.systems[index].keyreleased(key)
        end
    end

    w.resize = function(width, height)
        for index = 1, #w.systems do
            w.systems[index].resize(width, height)
        end
    end

    return w
end
