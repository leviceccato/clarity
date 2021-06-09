return function(name)
    local w = {}
    w.name = name
    w.systems = {}
    w.entities = {}

    w.load = function()
        for index = 1, #w.systems do
            w.systems[index].load()
        end
    end

    w.update = function()
        for index = 1, #w.systems do
            w.systems[index].update()
        end
    end

    w.draw = function()
        for index = 1, #w.systems do
            w.systems[index].draw()
        end
    end

    w.addEntity = function(entity)
        w.entities[entity.id] = entity
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
        local requiredComponent
        local component
        local hasComponent = false
        for systemIndex = 1, #w.systems do
            system = w.systems[systemIndex]
            for entityIndex = 1, #w.entities do
                entity = w.entities[entityIndex]
                for requiredComponentIndex = 1, #system.components do
                    requiredComponent = system.components[requiredComponentIndex]
                    for componentIndex = 1, #entity.components do
                        component = entity.components[componentIndex]
                        if component == requiredComponent then
                            hasComponent = true
                        end
                    end
                end
                if hasComponent then
                    system.addEntity(entity)
                    hasComponent = false
                end
            end
        end
    end

    return w
end
