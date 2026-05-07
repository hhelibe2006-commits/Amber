cmake_minimum_required(VERSION 4.2)
project(Amber)

set(CMAKE_CXX_STANDARD 14)

set(CMAKE_PREFIX_PATH "C:/Qt/6.11.0/mingw_64")

set(CMAKE_AUTOMOC ON)
set(CMAKE_AUTOUIC ON)
set(CMAKE_AUTORCC ON)

find_package(Qt6 REQUIRED COMPONENTS Widgets Core Gui)

add_executable(Amber
        src/main.cpp
        src/ui/main_ui.cpp
        src/ui/main_ui.h
        src/utils/centered_ui.h
        src/utils/set_window_size.h
        src/ui/settings_ui.cpp
        src/ui/settings_ui.h)

target_link_libraries(Amber PRIVATE Qt6::Widgets Qt6::Core Qt6::Gui)
